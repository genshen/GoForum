package account

import (
	"sync"
	"bytes"
	"text/template"
	"qiniupkg.com/x/log.v7"
	"github.com/astaxie/beego"
	"gensh.me/goforum/components/event"
	"gensh.me/goforum/components/utils"
)

const tpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>邮箱激活验证</title>
    <style>
        body{font-family:"Microsoft Yahei"}
        .button{ height: 40px;padding: 0 150px;border-radius: 4px;
            display:inline-flex;align-items:center;align-content: center;
            text-decoration: none;font-size: 15px;background-color: #0fb2fc; color: #FFFFFF;}
        .button:hover, .button:active{background-color: #139efc;}
        .span{margin-top: 10px; margin-bottom: 10px;}
        .line{margin-bottom:10px;}
        .line p{margin:5px 0 ;}
    </style>
</head>
<body>
    <div class="line">
        <p>亲爱的:&nbsp;<b>{{.Name}}</b></p>
        <p>欢迎加入clothesplus!</p>
        <p>请在48小时内点击下方链接以完成注册:(链接48小时有效)</p>
    </div>
    <a class="button" target="_blank" href="{{.Url}}">验证邮箱</a>
    <div class="span">
        <span>链接不起作用?复制下面的链接到浏览器</span>
    </div>
    <div class="line">
        <a href="{{.Url}}">{{.Url}}</a>
        <p>(提示:如果这不是你本人的注册行为,请不要点开链接,并删除本邮件)</p>
        <p>clothesplus团队</p>
        <p >（本邮件由系统自动发出，请勿回复.）</p>
    </div>
</body>
</html>`

var EnableEmail = false
var email_template *template.Template
var mutex sync.Mutex

func init() {
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatal("error to parse email template")
	}
	email_template = t
	EnableEmail = beego.AppConfig.DefaultBool("enable_email",false)
}

func OnAccountCreated(email string, username string, uid uint) {
	event.OnAccountCreated(email, username, uid)
	if(!EnableEmail){
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	var doc bytes.Buffer
	var data = map[string]string{"Name":username,"Url":"http://discus.kuluosi.com/account/?id=23&token=dsjiawUUYG76THJ7h7UU"}
	err := email_template.Execute(&doc, data)
	if err != nil{
		log.Println("error to execute email template.")
		return
	}
	s := doc.String()
	go utils.SendMail(email, "激活账号", &s)
}
