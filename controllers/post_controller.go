package controllers

import (
	"html/template"
	"strconv"
	"github.com/qiniu/api.v7/kodo"
	"./../models/forms"
	"./../models/m"
	form_check "./../verify/form"
	identify "./../verify/auth"
)

type PostController struct {
	BaseController
}

var post_rules = map[string]int{
	"View":   0,
	"CreateJump": identify.Login,
	"CreateMobile": identify.Login,
	"POST_CreateMobile": identify.Login,
	//UploadToken
}

func (this *PostController) getRules(action string) int {
	return post_rules[action]
}

func (this *PostController) CreateJump() {
	this.TplName = "post/create_jump.html"
}

func (this *PostController) CreateMobile() {
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "post/create_mobile.html"
}

func (this *PostController) POST_CreateMobile() {
	//this.Data["json"] = ""
	//this.ServeJSON()
	title := this.GetString("post_title")
	content := this.GetString("post_content")
	form := forms.PostCreateForm{Title:title, Content:content}
	if errs := form.Valid(); errs == nil {
		if id := form.Save(this.getUserId()); id != 0 {
			this.Redirect("/post/" + strconv.FormatInt(int64(id), 10), 302)
			//or use : beego.URLFor("",id)
			return
		}
	} else {
		//todo set error
		s := form_check.NewInstant(errs, map[string]string{"title":  title, "content": ""})
		this.Data["form_check"] = string(s)
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "post/create_mobile.html"
}

type QiNiuToken struct {
	Token string
}

func (this *PostController) UploadToken() {
	// 配置 AccessKey/SecretKey
	kodo.SetMac("xuCKJvuw7zg8qbgB1-LjcbqvS2H1O3SbTvO_zy7c", "YNutXAQzWvozFxPPfWbGz03XEHgOxdcADHmcrifQ")
	zone := 0
	c := kodo.New(zone, nil) // 创建一个 Client 对象
	policy := &kodo.PutPolicy{
		Scope:   "clothes-plus",
		Expires: 360,
	}
	up_token := c.MakeUptoken(policy)
	this.Data["json"] = &QiNiuToken{Token:up_token}
	this.ServeJSON()
}

func (this *PostController) View() {
	mPost := m.Posts{}
	mPost.GetPostById(this.Ctx.Input.Param(":id"))
	if mPost.ID != 0 {
		mUser := m.User{}
		mUser.GetUserById(mPost.Author) //todo query user profile
		//this.Data["data"] = mPost.Content
		this.Data["data"] = mUser.Password
		this.TplName = "post/view.html"
		return
	}
	this.Abort("404")
}