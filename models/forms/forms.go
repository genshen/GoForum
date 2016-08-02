package forms

import (
	"github.com/astaxie/beego/validation"
	"time"
	"gensh.me/goforum/middleware/event"
	"gensh.me/goforum/middleware/values"
	"gensh.me/goforum/models/database"
	"gensh.me/goforum/middleware/auth/security"
	"gensh.me/goforum/models/m"
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
	database.DB.Where("email = ? AND password_hash = ?", this.Username, security.Hash(this.Password)).First(&u)
	if u.ID == 0 {
		v.SetError("name", "用户名或密码错误")
		v.SetError("pass", "用户名或密码错误")
	} else {
		this.Username = u.Name  //use name to replace email or tel
	}
	return u.ID
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
	database.DB.Where("email = ?", this.Email).First(&u) //todo use NewRecord
	if u.ID != 0 {
		v.SetError("email", "该邮箱已经被使用")
	} else {
		user := m.User{Email:this.Email, Name:this.Nickname, Password:security.Hash(this.Password), Status:values.UNACTIVATED,
			Profile:m.Profile{Avatar:"/static/img/default.png"}}
		database.DB.Create(&user)
		this.UserID = u.ID
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
	//todo check user account first
	post := m.Posts{TopicID:1, AuthorID:userID, Title:this.Title, Summary:this.Summary,
		Content:this.Content, IsMobile:true, LastReplayAt:time.Now()};
	database.DB.Create(&post)
	return post.ID
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

func (this *CommentCreateForm)Create(user_id uint, username string) (result *SimpleJsonResponse) {
	p := m.Posts{}
	// use transaction
	tx := database.DB.Begin()
	if (p.Exist(this.PostID) ) {
		comment := m.Comment{PostID:this.PostID, Author:user_id, Content:this.Content}
		if err := tx.Create(&comment).Error; err != nil {
			tx.Rollback();
			result = &SimpleJsonResponse{Status:0, Addition:"添加回复失败,请重试!"}
			return
		}
		if err := tx.Model(&p).UpdateColumn("comment_count", p.CommentCount + 1).Error; err != nil {
			tx.Rollback()
			result = &SimpleJsonResponse{Status:0, Error:"添加回复失败,请重试!"}
			return
		}
		u := m.Profile{} //update profile attributes
		tx.Select("user_refer,comment_count").Where("user_refer = ?", user_id).First(&u)
		if err := tx.Table("profile").Where("user_refer = ?", user_id).UpdateColumn("comment_count", u.CommentCount + 1).Error; err != nil {
			tx.Rollback()
			result = &SimpleJsonResponse{Status:0, Error:"添加回复失败,请重试!"}
			return
		}
		result = &SimpleJsonResponse{Status:1, Addition:comment.ID}
		event.OnCommentSubmitted(&p, &comment, username)
	} else {
		result = &SimpleJsonResponse{Status:0, Error:"对应文章不存在"}
	}
	tx.Commit()
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
	database.DB.Where("status != ?", values.FREEZING).First(&u, this.PersonID)
	if u.ID == 0 {
		//todo use NewRecord
		result = &SimpleJsonResponse{Status:0, Error:"对应用户不存在"}
	} else {
		follow := m.Follow{}
		if database.DB.Where("follower_id = ? AND following_id = ?", my_id, this.PersonID).First(&follow);
		follow.FollowerID == 0 && follow.FollowingID == 0 {
			database.DB.Create(&m.Follow{FollowerID:my_id, FollowingID:this.PersonID})
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
	feedback := m.Feedback{UserID:uid, Type:this.Type, Content:this.Feedback, Contact:this.Contact}
	database.DB.Create(&feedback)
	return nil
}
