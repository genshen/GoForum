package event

import (
	"fmt"
	"log"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego"
	"net/textproto"
	"../values"
	"../../models/m"
	"./../../models/database"
)

func OnAccountCreated(email string, username string, uid uint) {
	go sendMail(email, "激活账号", "Text Message")
}

//todo: none
func OnPostCreated() {

}

//Posts:id,title,comment_count
func OnCommentSubmitted(post *m.Posts, comment *m.Comment, username string) {
	var message = m.PostMessage{RelatedUsername:username, PostID:post.ID, PostTitle:post.Title,
		BaseMessage:m.BaseMessage{UserID:post.AuthorID, RelatedID:comment.Author, SubjectType:values.POST_COMMENT} }
	database.DB.Create(&message)
	//多次create,message 可以复用
}

/*id userId;id_r:related_id;name:*/
func OnFollowed(id uint, id_r uint, name string) {
	var notify = m.Notification{Data:"{\"username\":\"" + name + "\"}",
		BaseMessage:m.BaseMessage{UserID:id, RelatedID:id_r, SubjectType:values.FOLLOW_ADD}}
	database.DB.Create(&notify)
}

func OnUnFollowed() {

}

func sendMail(to string, subject string, content string) {
	port, _ := beego.AppConfig.Int("mail_port")
	var mail = utils.Email{
		Username:beego.AppConfig.String("mail_user"),
		Password:beego.AppConfig.String("mail_password"),
		Host:beego.AppConfig.String("mail_host"),
		Port:port,
		From:beego.AppConfig.String("mail_from"),
	}
	mail.Headers = textproto.MIMEHeader{}
	mail.To = []string{to}
	mail.Subject = subject
	mail.HTML = content
	err := mail.Send()
	if err != nil {
		log.Println("send to:" + to)
		log.Println(fmt.Sprintf("Send Mail Error:%v", err))
	}
}
