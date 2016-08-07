package controllers

import (
	"strconv"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/components/context/comments"
	identify "gensh.me/goforum/models/values"
)

type CommentController struct {
	BaseController
}

var comment_rules = map[string]int{
	"View":   0,
	"CommentAdd":identify.LoginJSON,
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
		mComments = append(mComments, Comment{Content:comment.Content, ID:comment.Id});
	}
	this.Data["json"] = &mComments
	this.ServeJSON()
}

//post only
// 0 for article deleted; 1 for success,3 for no login;
func (this *CommentController) CommentAdd() {
	var result *utils.SimpleJsonResponse
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))  //string to int
	content := this.GetString("content")
	ccf := comments.CommentCreateForm{PostID:uint(id), Content:content}
	result = ccf.Create(this.getUserId(),this.getUsername())
	this.Data["json"] = result
	this.ServeJSON()
}
