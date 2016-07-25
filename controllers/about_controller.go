package controllers

import (
	//"../models/forms"
)
import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
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

func (this *AboutController) Feedback() {
	this.TplName = "about/feedback.html"
}

func (this *AboutController)POST_Feedback(){
//forms.PostResult{}
}
