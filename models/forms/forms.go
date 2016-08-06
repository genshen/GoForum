package forms

import (
	"github.com/astaxie/beego/validation"
	"time"
	"gensh.me/goforum/middleware/event"
	"gensh.me/goforum/middleware/values"
	"gensh.me/goforum/models/database"
	"gensh.me/goforum/middleware/auth/security"
	"gensh.me/goforum/models/m"
	"github.com/astaxie/beego/orm"
)

type SignInForm struct {
	Username string
	Password string
}

type SignUpForm struct {
	UserID   uint
	Email    string
	Nickname string
	Password string
}

type PostCreateForm struct {
	Title   string
	Summary string
	Content string
}

func (this *SignInForm)SignInVerify() ([]*validation.Error, uint) {
	valid := validation.Validation{}
	valid.Required(this.Username, "name").Message("用户名不能为空")
	valid.Required(this.Password, "pass").Message("密码不能为空")
	if valid.HasErrors() {
		return valid.Errors, 0
	}

	id := this.validPassword(&valid); //验证密码
	if valid.HasErrors() {
		return valid.Errors, 0
	}
	return nil, id
}

func (this *SignInForm) validPassword(v *validation.Validation) uint {
	u := m.User{};
	database.O.QueryTable("user").Filter("email",this.Username).
	Filter("password_hash", security.Hash(this.Password)).One(&u)
	//RW database.DB.Where("email = ? AND password_hash = ?", this.Username,.First(&u)
	if u.Id == 0 {
		v.SetError("name", "用户名或密码错误")
		v.SetError("pass", "用户名或密码错误")
	} else {
		this.Username = u.Name  //use name to replace email or tel
	}
	return u.Id
}

func (this *SignUpForm)Valid() ([]*validation.Error) {
	valid := validation.Validation{}
	valid.Required(this.Email, "email").Message("邮箱不能为空")
	valid.Email(this.Email, "email").Message("邮箱格式不正确")
	valid.Required(this.Nickname, "nickname").Message("昵称不能为空")
	valid.Required(this.Password, "pass").Message("密码不能为空")
	valid.MinSize(this.Password, 6, "pass").Message("密码长度不能小于6")
	if valid.HasErrors() {
		return valid.Errors
	}
	this.validOrSave(&valid);
	if valid.HasErrors() {
		return valid.Errors
	}
	return nil
}

func (this *SignUpForm)validOrSave(v *validation.Validation) {
	u := m.User{};
	database.O.QueryTable("user").Filter("email",this.Email).One(&u)
	//database.DB.Where("email = ?", this.Email).First(&u) //todo use NewRecord
	if u.Id != 0 {
		v.SetError("email", "该邮箱已经被使用")
	} else {
		user := m.User{Email:this.Email, Name:this.Nickname, Password:security.Hash(this.Password), Status:values.UNACTIVATED,
			Profile:&m.Profile{Avatar:"/static/img/default.png"}} //todo Profile
		database.O.Insert(&user)//RW
		this.UserID = u.Id
	}
}

func (this *PostCreateForm)Valid() ([]*validation.Error) {
	valid := validation.Validation{}
	valid.Required(this.Title, "title").Message("标题不能为空")
	valid.Required(this.Content, "conten").Message("内容不能为空")
	valid.Required(this.Summary, "summary").Message("摘要不能为空")
	valid.MaxSize(this.Summary, 255, "summary").Message("摘要不超过255")
	if valid.HasErrors() {
		return valid.Errors
	}
	return nil
}

func (this *PostCreateForm)Save(userID uint) uint {
	//todo check user account first  AND save TopicId:1, AuthorId:userID,
	post := m.Posts{Title:this.Title, Summary:this.Summary,
		Content:this.Content, IsMobile:true, LastReplayAt:time.Now()};
	database.O.Insert(&post)//RW
	return post.Id
}

type CommentCreateForm struct {
	PostID  uint
	Content string
}

type SimpleJsonResponse struct {
	Status   int
	Error    interface{}
	Addition interface{}
}

//RW all
func (this *CommentCreateForm)Create(user_id uint, username string) (result *SimpleJsonResponse) {
	p := m.Posts{}
	// use transaction
	o := orm.NewOrm()
	o.Begin()
	//tx := database.DB.Begin()
	if (p.Exist(this.PostID) ) {
		comment := m.Comment{PostId:this.PostID, Author:user_id, Content:this.Content}
		if _,err := o.Insert(&comment); err != nil {
			//tx.Rollback();
			//fmt.Println(err,err ==nil )
			err = o.Rollback()
			result = &SimpleJsonResponse{Status:0, Addition:"添加回复失败,请重试!"}
			return
		}
		p.CommentCount++;
		if _,err := o.Update(&p); err != nil {
			//tx.Rollback()
			err = o.Rollback()
			result = &SimpleJsonResponse{Status:0, Error:"添加回复失败,请重试!"}
			return
		}
		u := m.Profile{} //update profile attributes
		o.QueryTable("profile").Filter("user_refer",user_id).One(&u,"user_refer","comment_count")
		//RW Select("user_refer,comment_count").Where("user_refer = ?", user_id).First(&u)
		u.CommentCount++
		if _,err := o.Update(&u); err != nil {
			//tx.Rollback()
			err = o.Rollback()
			result = &SimpleJsonResponse{Status:0, Error:"添加回复失败,请重试!"}
			return
		}
		result = &SimpleJsonResponse{Status:1, Addition:comment.Id}
		event.OnCommentSubmitted(&p, &comment, username)
	} else {
		result = &SimpleJsonResponse{Status:0, Error:"对应文章不存在"}
	}
	o.Commit()
	return
}

type FollowAddForm struct {
	PersonID uint
}

func (this *FollowAddForm)Add(my_id uint) (result *SimpleJsonResponse, isAdded bool) {
	isAdded = false
	if (my_id == this.PersonID) {
		result = &SimpleJsonResponse{Status:0, Error:"不能关注自己本人哟~"}
		return
	}
	u := m.User{}
	database.O.QueryTable("user").Filter("id", this.PersonID).Filter("status", values.FREEZING).One(&u) //todo FREEZING
	if u.Id == 0 {
		//todo use NewRecord
		result = &SimpleJsonResponse{Status:0, Error:"对应用户不存在"}
	} else {
		follow := m.Follow{}
		//if database.DB.Where("follower_id = ? AND following_id = ?", my_id, this.PersonID).First(&follow);
		//follow.FollowerID == 0 && follow.FollowingID == 0 {
		if database.O.QueryTable("follows").Filter("follower_id",my_id).Filter("following_id",this.PersonID).One(&follow);
		follow.FollowerID == 0 && follow.FollowingID == 0 {
			database.O.Insert(&m.Follow{FollowerID:my_id, FollowingID:this.PersonID})//RW
			result = &SimpleJsonResponse{Status:1, Addition:this.PersonID}
			isAdded = true
		} else {
			result = &SimpleJsonResponse{Status:0, Error:"已经关注了改用户"}
		}
	}
	return
}

type FeedbackForm struct {
	Type     int8
	Feedback string
	Captcha  bool
	Contact  string
}

func (this *FeedbackForm)Valid(uid uint) ([]*validation.Error) {
	v := validation.Validation{}
	if !this.Captcha {
		v.SetError("captcha", "验证码错误")
		return v.Errors
	}
	//todo 联系方式验证
	v.Required(this.Feedback, "feedback").Message("反馈意见不能为空")
	v.MaxSize(this.Feedback, 512, "feedback").Message("反馈意见字数不能超过255")
	if this.Type > 2 || this.Type < 0 {
		v.SetError("type", "反馈类型错误")
	}
	if v.HasErrors() {
		return v.Errors
	}
	feedback := m.Feedback{UserId:uid, Type:this.Type, Content:this.Feedback, Contact:this.Contact}
	database.O.Insert(&feedback) //RW
	return nil
}
