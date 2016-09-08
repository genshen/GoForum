package utils

import (
	"fmt"
	"log"
	"net/textproto"
	"github.com/astaxie/beego"
	_util "github.com/astaxie/beego/utils"
)

func SendMail(to string, subject string, content *string) {
	port, _ := beego.AppConfig.Int("mail_port")
	var mail = _util.Email{
		Username:beego.AppConfig.String("mail_user"),
		Password:beego.AppConfig.String("mail_password"),
		Host:beego.AppConfig.String("mail_host"),
		Port:port,
		From:beego.AppConfig.String("mail_from"),
	}
	mail.Headers = textproto.MIMEHeader{}
	mail.To = []string{to}
	mail.Subject = subject
	mail.HTML = *content
	err := mail.Send()
	if err != nil {
		log.Println("send to:" + to)
		log.Println(fmt.Sprintf("Send Mail Error:%v", err))
	}
}

