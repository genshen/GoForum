package controllers

import (
	"../models/m"
	"../models/forms"
	"strconv"
)

type CommentController struct {
	BaseController
}

var comment_rules = map[string]int{
	"View":   0,
}

func (this *CommentController) getRules(action string) int {
	return comment_rules[action]
}

type Comment struct {
	Content string
	ID      uint
}

/*show 20 comments ordered by submit time of one post
id post id;  start:comment offset  */
func (this *CommentController) Comment() {
	//todo ([]m.Comment) to []Comment
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	offset, _ := strconv.Atoi(this.Ctx.Input.Param(":start"))
	dbComments := m.LoadComments(id, offset)
	mComments := make([]Comment, 0, len(dbComments))  //m.Comments to Comments
	for _, comment := range dbComments {
		mComments = append(mComments, Comment{Content:comment.Content, ID:comment.ID});
	}
	this.Data["json"] = &mComments
	this.ServeJSON()
}

//post only
// 0 for no login;2 for article deleted; 1 feo success
func (this *CommentController) CommentAdd() {
	var result *forms.PostResult
	if !this.IsUserLogin() {
		result = &forms.PostResult{Status:0, Error:"用户未登录"}
	} else {
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))  //string to int
		content := this.GetString("content")
		ccf :=  forms.CommentCreateForm{PostID:uint(id),Content:content}
		result  = ccf.Create(this.getUserId())
	}
	this.Data["json"] = result
	this.ServeJSON()
}