package controllers

import (
	"html/template"
	"strconv"
	"github.com/qiniu/api.v7/kodo"
	"encoding/json"
	"gensh.me/goforum/models/forms"
	"gensh.me/goforum/middleware/event"
	"gensh.me/goforum/models/m"
	form_check "gensh.me/goforum/middleware/form"
	identify "gensh.me/goforum/middleware/values"
)

type PostController struct {
	BaseController
}

type QiNiuToken struct {
	Token string
}

var post_rules = map[string]int{
	"View":   0,
	"CreateJump": identify.Login | identify.JumpBack,
	"CreateMobile": identify.Login | identify.JumpBack,
	"POST_CreateMobile": identify.Login | identify.JumpBack,
	"UploadToken":identify.LoginJSON,
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
	summary := this.GetString("post_summary")
	form := forms.PostCreateForm{Title:title, Summary:summary, Content:content}
	if errs := form.Valid(); errs == nil {
		if id := form.Save(this.getUserId()); id != 0 {
			this.Redirect("/post/" + strconv.FormatInt(int64(id), 10), 302)
			event.OnPostCreated()
			//or use : beego.URLFor("",id)
			return
		}
	} else {
		//todo set error
		s := form_check.NewInstantToByte(errs, map[string]string{"title":  title, "summary":summary, "content": content})
		this.Data["form_check"] = string(s)
	}
	this.Data["xsrf_token"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "post/create_mobile.html"
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
		mPostDetail := PostDetail{}
		mPostDetail.NewInstant(&mPost)
		mUser := m.User{}
		mUser.GetUserById(mPost.AuthorID) //todo query user profile,do not show Author(in mPost) information
		person := Person{ID:mUser.ID, Name:mUser.Name, Avatar:mUser.Profile.Avatar}
		data := PostView{IsLogin:this.IsUserLogin(), Post:mPostDetail, Author:person}
		json, err := json.Marshal(data)
		if err == nil {
			this.Data["data"] = string(json)
			this.TplName = "post/view.html"
			return
		}
	}
	this.Abort("404")
}