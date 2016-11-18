package controllers

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"gensh.me/goforum/components/context/about"
	"gensh.me/goforum/components/utils"
)

var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/temp/captcha/", store)
}

type AboutController struct {
	BaseController
}

var about_rules = map[string]int{
}

func (this *AboutController) getRules(action string) int {
	return about_rules[action]
}

func (this *AboutController) Index() {
	this.TplName = "about/index.html"
}

func (this *AboutController) Agreement() {
	this.TplName = "about/agreement.html"
}

func (this *AboutController) Feedback() {
	this.TplName = "about/feedback.html"
}

func (this *AboutController)POST_Feedback() {
	post_result := utils.SimpleJsonResponse{Status:1}
	captcha := cpt.VerifyReq(this.Ctx.Request)
	feedback := this.GetString("feedback")
	contact := this.GetString("contact")
	ty, ty_err := this.GetInt8("type")
	if ty_err != nil {
		ty = -1;
	}

	feedback_form := about.FeedbackForm{Type:ty, Captcha:captcha, Feedback:feedback, Contact:contact}
	if errs := feedback_form.Valid(this.getUserId()); errs == nil {
		post_result.Addition = "tokens"
	} else {
		post_result.Status = 0
		s := utils.NewInstant(errs, map[string]string{"type":"", "feedback":"", "captcha": ""})
		post_result.Error = &s
	}
	this.Data["json"] = &post_result;
	this.ServeJSON()
}
