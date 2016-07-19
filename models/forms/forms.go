package forms

import (
	"../database"
	"../../verify/auth/security"
	"../m"
	"github.com/astaxie/beego/validation"
	"time"
)

type SignInForm struct {
	Username string
	Password string
}

type SignUpForm struct {
	Email    string
	Nickname string
	Password string
}

type PostCreateForm struct {
	Title   string
	Content string
}

func (this *SignInForm)LoginVerify() ([]*validation.Error, uint) {
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

func (this *SignUpForm)Valid() ([]*validation.Error, bool) {
	valid := validation.Validation{}
	valid.Required(this.Email, "email").Message("邮箱不能为空")
	valid.Email(this.Email, "email").Message("邮箱格式不正确")
	valid.Required(this.Nickname, "nickname").Message("昵称不能为空")
	valid.Required(this.Password, "pass").Message("密码不能为空")
	valid.MinSize(this.Password, 6, "pass").Message("密码长度不能小于6")
	if valid.HasErrors() {
		return valid.Errors, false
	}
	this.validOrSave(&valid);
	if valid.HasErrors() {
		return valid.Errors, false
	}
	return nil, true
}

func (this *SignUpForm)validOrSave(v *validation.Validation) {
	u := m.User{};
	database.DB.Where("email = ?", this.Email).First(&u)
	if u.ID != 0 {
		v.SetError("email", "该邮箱已经被使用")
	} else {
		user := m.User{Email:this.Email, Name:this.Nickname, Password:security.Hash(this.Password), Status:m.STATUS_ACTIVE,
			Profile:m.Profile{Avatar:"/static/img/default.png"}}
		database.DB.Create(&user)
	}
}

func (this *PostCreateForm)Valid() ([]*validation.Error) {
	valid := validation.Validation{}
	valid.Required(this.Title, "title").Message("标题不能为空")
	valid.Required(this.Content, "conten").Message("内容不能为空")
	if valid.HasErrors() {
		return valid.Errors
	}
	return nil
}

func (this *PostCreateForm)Save(userID uint) uint {
	//todo check user account first
	post := m.Posts{TopicID:1, AuthorID:userID, Title:this.Title, Content:this.Content, IsMobile:true, LastReplayAt:time.Now()};
	database.DB.Create(&post)
	return post.ID
}

type CommentCreateForm struct {
	PostID  uint
	Content string
}

type CommentAddResult struct {
	Status   int
	Error    string
	Addition interface{}
}

func (this *CommentCreateForm)Create(user_id uint) (result *CommentAddResult) {
	p := m.Posts{}
	// use transaction
	tx := database.DB.Begin()
	if (p.Exist(this.PostID) ) {
		comment := m.Comment{PostID:this.PostID, Author:user_id, Content:this.Content}
		if err := database.DB.Create(&comment).Error; err != nil {
			tx.Rollback();
			result = &CommentAddResult{Status:2, Addition:0}
			return
		}
		if err := database.DB.Model(&p).UpdateColumn("CommentCount", p.CommentCount + 1).Error; err != nil {
			tx.Rollback();
			result = &CommentAddResult{Status:2, Addition:0}
			return
		}
		result = &CommentAddResult{Status:1, Addition:comment.ID}
	} else {
		result = &CommentAddResult{Status:3, Error:"对应文章不存在"}
	}
	tx.Commit()
	return
}