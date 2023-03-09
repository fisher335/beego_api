package routers

import (
	"beego_api/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	ns := beego.NewNamespace("/user",
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/register", &controllers.UserController{}, "post:Post"),
		beego.NSRouter("/del", &controllers.UserController{}, "post:Delete"),
		beego.NSRouter("/list", &controllers.UserController{}, "post:GetAll"),
	)
	beego.AddNamespace(ns)

}
