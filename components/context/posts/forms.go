package posts

import (
	"time"
	"fmt"
	"github.com/astaxie/beego/validation"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/models/database"
)

type PostCreateForm struct {
	PostID  uint
	Title   string
	Summary string
	Content string
}

func (this *PostCreateForm)Valid(userId uint) ([]*validation.Error) {
	valid := validation.Validation{}
	valid.Required(this.Title, "title").Message("标题不能为空")
	valid.Required(this.Content, "conten").Message("内容不能为空")
	valid.Required(this.Summary, "summary").Message("摘要不能为空")
	valid.MaxSize(this.Summary, 255, "summary").Message("摘要不超过255")
	if valid.HasErrors() {
		return valid.Errors
	}
	this.Save(&valid, userId)
	return nil
}

func (this *PostCreateForm)Save(valid *validation.Validation, userID uint) {
	//todo check user account first  AND save TopicId:1, AuthorId:userID,

	post := m.Posts{Title:this.Title, Summary:this.Summary, Topic:&m.Topic{Id:1}, Author:&m.Profile{Id:userID}, //todo Topic
		Content:this.Content, IsMobile:true, Visible:true, LastReplayAt:time.Now()};
	id, err := database.O.Insert(&post) //RW
	if err == nil {
		this.PostID = uint(id)
	} else {
		fmt.Println(err)
		valid.SetError("conten", "创建帖子的过程中发生了一些问题,请重试")
	}
}

