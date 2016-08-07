package about

import (
	"github.com/astaxie/beego/validation"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
)

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

