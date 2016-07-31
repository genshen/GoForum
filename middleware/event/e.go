package event

import (
	"fmt"
	"log"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego"
	"net/textproto"
	"../../models/m"
	"./../../models/database"
)

func OnAccountCreated(email string, username string, uid uint) {
	go sendMail(email, "激活账号", "Text Message")
}

//todo: none
func OnPostCreated() {

}

func OnCommentSubmitted() {

}

/*id userId;id_r:related_id;name:*/
func OnFollowed(id uint, id_r uint, name string) {
	var notify = m.Notification{UserID:id, RelatedID:id_r, Title:name + "关注了你", Subject:"关注", SubjectType:1}
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