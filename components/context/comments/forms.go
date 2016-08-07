package comments

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gensh.me/goforum/models/m"
	"gensh.me/goforum/components/utils"
	"gensh.me/goforum/components/event"
)

type CommentCreateForm struct {
	PostID  uint
	Content string
}

//RW all
func (this *CommentCreateForm)Create(user_id uint, username string) (result *utils.SimpleJsonResponse) {
	p := m.Posts{}
	// use transaction
	o := orm.NewOrm()
	o.Begin()
	//tx := database.DB.Begin()
	if (p.Exist(this.PostID) ) {
		comment := m.Comment{PostId:this.PostID, Author:user_id, Content:this.Content}
		if id, err := o.Insert(&comment); err != nil {
			err = o.Rollback()
			result = &utils.SimpleJsonResponse{Status:0, Addition:"添加回复失败,请重试!"}
			return
		} else {
			comment.Id = uint(id)
		}

		num, err := o.QueryTable("posts").Filter("id", this.PostID).Update(orm.Params{
			"comment_count": orm.ColValue(orm.ColAdd, 1),
		})
		if num == 0 || err != nil {
			fmt.Println(err)
			err = o.Rollback()
			result = &utils.SimpleJsonResponse{Status:0, Error:"添加回复失败,请重试!"}
			return
		}

		num, err = o.QueryTable("profile").Filter("id", user_id).Update(orm.Params{
			"comment_count": orm.ColValue(orm.ColAdd, 1),
		})
		if num == 0 || err != nil {
			//tx.Rollback()
			err = o.Rollback()
			result = &utils.SimpleJsonResponse{Status:0, Error:"添加回复失败,请重试!"}
			return
		}
		result = &utils.SimpleJsonResponse{Status:1, Addition:comment.Id}
		event.OnCommentSubmitted(&p, &comment, username)
	} else {
		result = &utils.SimpleJsonResponse{Status:0, Error:"对应文章不存在"}
	}
	o.Commit()
	return
}
