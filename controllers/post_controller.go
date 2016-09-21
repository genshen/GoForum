package controllers

import (
	"strconv"
	"html/template"
	"encoding/json"
	"github.com/qiniu/api.v7/kodo"
	"gensh.me/goforum/components/event"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/context/posts"
	"gensh.me/goforum/components/utils"
)

type PostController struct {
	BaseController
}

type QiNiuToken struct {
	Token string
}

var post_rules = map[string]int{
	"View":   0,
	"CreateJump": utils.Login | utils.JumpBack,
	"CreateMobile": utils.Login | utils.JumpBack,
	"POST_CreateMobile": utils.Login | utils.JumpBack,
	"UploadToken":utils.LoginJSON,
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
	form := posts.PostCreateForm{Title:title, Summary:summary, Content:content}
	if errs := form.Valid(this.getUserId()); errs == nil {
		//if id := form.Save(this.getUserId()); id != 0 {
			this.Redirect("/post/" + strconv.FormatInt(int64(form.PostID), 10), 302)
			event.OnPostCreated() //&form
			//or use : beego.URLFor("",id)
			return
		//}
	} else {
		//todo set error
		s := utils.NewInstantToByte(errs, map[string]string{"title":  title, "summary":summary, "content": content})
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
	err := mPost.GetPostById(this.Ctx.Input.Param(":id"))
	if err == nil {
		mPostDetail :=  posts.PostDetail{}
		mPostDetail.NewInstant(&mPost)
		person := utils.Person{ID:mPost.Author.Id, Name:mPost.Author.Name, Avatar:mPost.Author.Avatar}
		data :=  posts.PostView{IsLogin:this.IsUserLogin(), Post:mPostDetail, Author:person}
		j, err := json.Marshal(data)
		if err == nil {
			this.Data["data"] = string(j)
			this.TplName = "post/view.html"
			return
		}
	}
	this.Abort("404")
}