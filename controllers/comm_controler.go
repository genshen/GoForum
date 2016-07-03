package controllers

import (
	"../models/m"
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

type Comments struct {
	Size     int
	Comments []Comment
}

/*show 20 comment ordered by submit time of one post
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
	comments := Comments{Size:len(mComments), Comments:mComments}
	this.Data["json"] = &comments
	this.ServeJSON()
}

type CommentAddResult struct {
	Status   int
	Error    string
	Addition interface{}
}
//post only
// 0 for no login;2 for article deleted; 1 feo success
func (this *CommentController) CommentAdd() {
	var result CommentAddResult
	if !this.IsUserLogin() {
		result = CommentAddResult{Status:0, Error:"用户未登录"}
	} else {
		mPost := m.Posts{}
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))  //string to int
		post_id := uint(id);
		content := this.GetString("content")
		if (mPost.Exist(post_id) ) {
			//save it
			comment := m.Comment{PostID:post_id, Author:this.getUserId(), Content:content}
			if comment.Create() {
				result = CommentAddResult{Status:1, Addition:0}
			} else {
				result = CommentAddResult{Status:2, Addition:0}
			}
		} else {
			result = CommentAddResult{Status:3, Error:"对应文章不存在"}
		}
	}
	this.Data["json"] = &result
	this.ServeJSON()
}