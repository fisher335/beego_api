package routers

import (
	"beego_api/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	ns := beego.NewNamespace("/user",
		beego.NSRouter("/login", &controllers.UserController{}, "*:Login"),
	)
	beego.AddNamespace(ns)

}
