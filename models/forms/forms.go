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
	Password string
}

type PostCreateForm struct {
	Title   string
	Content string
}

func (this *SignInForm)LoginVerify() ([]*validation.Error, bool, uint) {
	valid := validation.Validation{}
	valid.Required(this.Username, "name").Message("用户名不能为空")
	valid.Required(this.Password, "pass").Message("密码不能为空")
	if valid.HasErrors() {
		return valid.Errors, false, 0
	}

	id := this.validPassword(&valid); //验证密码
	if valid.HasErrors() {
		return valid.Errors, false, 0
	}
	return nil, true, id
}

func (this *SignInForm) validPassword(v *validation.Validation) uint {
	u := m.User{};
	database.DB.Where("email = ? AND password_hash = ?", this.Username, security.Hash(this.Password)).First(&u)
	if u.ID == 0 {
		v.SetError("name", "用户名或密码错误")
		v.SetError("pass", "用户名或密码错误")
	}
	return u.ID
}

func (this *SignUpForm)Valid() ([]*validation.Error, bool) {
	valid := validation.Validation{}
	valid.Required(this.Email, "email").Message("邮箱不能为空")
	valid.Email(this.Email, "email").Message("邮箱格式不正确")
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
	database.DB.Where("email = ?", this.Password).First(&u)
	if u.ID != 0 {
		v.SetError("email", "该邮箱已经被使用")
	} else {
		user := m.User{Email:this.Email, Name:"", Password:security.Hash(this.Password), Status:m.STATUS_ACTIVE}
		database.DB.Create(&user)
	}
}

func (this *PostCreateForm)Valid() ([]*validation.Error){
	valid := validation.Validation{}
	valid.Required(this.Title, "title").Message("标题不能为空")
	valid.Required(this.Content, "conten").Message("内容不能为空")
	if valid.HasErrors() {
		return valid.Errors
	}
	return nil
}

func (this *PostCreateForm)Save(userID uint) uint{
	//todo check user account first
	post := m.Posts{Topic:1, Author:userID, Title:this.Title, Content:this.Content, IsMobile:true, LastReplayAt:time.Now()};
	database.DB.Create(&post)
	return post.ID
}