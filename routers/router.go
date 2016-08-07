package routers

import (
	"gensh.me/goforum/controllers"
	"github.com/astaxie/beego"
)

func init() {
	/*OK */beego.Router("/", &controllers.HomeController{})
	/*OK */beego.Router("home/swipe", &controllers.HomeController{}, "get:LoadSwipe") /*OK */
	/*OK */beego.Router("home/hot/:start([0-9]+)", &controllers.HomeController{}, "get:Hot")
	/*OK */beego.Router("home/category", &controllers.HomeController{}, "get:Category")
	/*OK */beego.Router("home/me", &controllers.HomeController{}, "get:Me")

	/*OK */	beego.Router("account/signup", &controllers.AccountController{}, "get:SignUp;post:POST_SignUp")
	/*OK */	beego.Router("account/signup/success", &controllers.AccountController{}, "get:SignUpSuccess")
	/*OK */	beego.Router("account/signin", &controllers.AccountController{}, "get:SignIn;post:POST_SignIn")
	/*OK */	beego.Router("account/signout", &controllers.AccountController{}, "get:SignOut")

	/*OK */	beego.Router("person/:uid([0-9]+)", &controllers.ProfileController{}, "get:Person")
	beego.Router("profile/following", &controllers.ProfileController{}, "get:Follow")
	beego.Router("profile/followed", &controllers.ProfileController{}, "get:Follow")
	beego.Router("profile/follow/add", &controllers.ProfileController{}, "post:FollowAdd")
	/*OK */beego.Router("profile/collection", &controllers.ProfileController{}, "get:Collection")

	/*OK */beego.Router("message/messages", &controllers.MessageController{}, "get:Messages")
	beego.Router("message/notifications", &controllers.MessageController{}, "get:Notifications")

	/*OK */beego.Router("topic/:slug([\\w]+)", &controllers.TopicController{}, "get:Slug")  //topic and tag

	/*OK */beego.Router("post/:id([0-9]+)", &controllers.PostController{}, "get:View")
	/*OK */beego.Router("post/create/jump", &controllers.PostController{}, "get:CreateJump")
	/*OK */beego.Router("post/create/m", &controllers.PostController{}, "get:CreateMobile;post:POST_CreateMobile")
	/*OK */beego.Router("post/create/upload_token", &controllers.PostController{}, "get:UploadToken")

	/*OK */beego.Router("comment/:id([0-9]+)/:start([0-9]+)", &controllers.CommentController{}, "get:Comment")
	/*OK */beego.Router("comment/add/:id([0-9]+)", &controllers.CommentController{}, "post:CommentAdd")
	//beego.AutoRouter(&controllers.UserController{})

	/*about */
	/*OK */beego.Router("about", &controllers.AboutController{}, "get:Index")
	/*OK */beego.Router("about/feedback", &controllers.AboutController{}, "get:Feedback;post:POST_Feedback")
}