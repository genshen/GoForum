package routers

import (
	"gensh.me/goforum/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("home/swipe", &controllers.HomeController{}, "get:LoadSwipe")
	beego.Router("home/hot/:start([0-9]+)", &controllers.HomeController{}, "get:Hot")
	beego.Router("home/category", &controllers.HomeController{}, "get:Category")
	beego.Router("home/me", &controllers.HomeController{}, "get:Me")

	beego.Router("account/signup", &controllers.AccountController{}, "get:SignUp;post:POST_SignUp")
	beego.Router("account/signup/success", &controllers.AccountController{}, "get:SignUpSuccess")
	beego.Router("account/signin", &controllers.AccountController{}, "get:SignIn;post:POST_SignIn")
		beego.Router("account/signout", &controllers.AccountController{}, "get:SignOut")

	beego.Router("person/:uid([0-9]+)", &controllers.ProfileController{}, "get:Person")
	beego.Router("profile/following", &controllers.ProfileController{}, "get:Follow")
	beego.Router("profile/followed", &controllers.ProfileController{}, "get:Follow")
	beego.Router("profile/follow/add", &controllers.ProfileController{}, "post:FollowAdd")
	beego.Router("profile/collection", &controllers.ProfileController{}, "get:Collection")

	beego.Router("message/messages", &controllers.MessageController{}, "get:Messages")
	beego.Router("message/notifications", &controllers.MessageController{}, "get:Notifications")

	beego.Router("topic/:slug([\\w]+)", &controllers.TopicController{}, "get:Slug")  //topic and tag

	beego.Router("post/:id([0-9]+)", &controllers.PostController{}, "get:View")
	beego.Router("post/create/jump", &controllers.PostController{}, "get:CreateJump")
	beego.Router("post/create/m", &controllers.PostController{}, "get:CreateMobile;post:POST_CreateMobile")
	beego.Router("post/create/upload_token", &controllers.PostController{}, "get:UploadToken")

	beego.Router("comment/:id([0-9]+)/:start([0-9]+)", &controllers.CommentController{}, "get:Comment")
	beego.Router("comment/add/:id([0-9]+)", &controllers.CommentController{}, "post:CommentAdd")
	//beego.AutoRouter(&controllers.UserController{})

	/*about */
	beego.Router("about", &controllers.AboutController{}, "get:Index")
	beego.Router("about/agreement", &controllers.AboutController{}, "get:Agreement")
	beego.Router("about/feedback", &controllers.AboutController{}, "get:Feedback;post:POST_Feedback")
}