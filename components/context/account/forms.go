package account

import (
	"github.com/astaxie/beego/validation"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/auth"
	"gensh.me/goforum/models/values"
	"gensh.me/goforum/models/database"
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
	database.O.QueryTable("user").Filter("email", this.Username).
	Filter("password_hash", auth.Hash(this.Password)).Limit(1).One(&u)
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
	valid.MaxSize(this.Nickname, 64, "nickname").Message("昵称长度不能超过64")
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
	database.O.QueryTable("user").Filter("email", this.Email).Limit(1).One(&u)
	//database.DB.Where("email = ?", this.Email).First(&u) //todo use NewRecord
	if u.Id != 0 {
		v.SetError("email", "该邮箱已经被使用")
	} else {
		//todo 事务
		user := m.User{Email:this.Email, Name:this.Nickname, Password:auth.Hash(this.Password), Status:values.UNACTIVATED}
		if id, err := database.O.Insert(&user); err == nil {
			profile := m.Profile{Id:uint(id), UserRefer:&m.User{Id:uint(id)}, Name:this.Nickname, Avatar:"/static/img/default.png"}
			if _, err := database.O.Insert(&profile); err == nil {
				return
			}
		}
		v.SetError("email", "创建账户失败,请重试")
	}
}
