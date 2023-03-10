package routers

import (
	"beego_api/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	userRouter := beego.NewNamespace("/user",
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/register", &controllers.UserController{}, "post:Post"),
		beego.NSRouter("/del", &controllers.UserController{}, "post:Delete"),
		beego.NSRouter("/list", &controllers.UserController{}, "post:GetAll"),
	)
	scanRouter := beego.NewNamespace("/scan",
		beego.NSRouter("/scan", &controllers.UserController{}, "post:Login"),
	)
	devRouter := beego.NewNamespace("/dev",
		beego.NSRouter("/list", &controllers.DeviceController{}, "post:GetAll"),
	)
	beego.AddNamespace(userRouter)
	beego.AddNamespace(scanRouter)
	beego.AddNamespace(devRouter)

}
