package common

import (
	"beego_api/services"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"strings"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	beego.InsertFilter("/*", beego.BeforeRouter, func(context *context.Context) {
		//此处可以校验一下ip，设备等
		if strings.Index(context.Request.RequestURI, "/user/login") >= 0 || strings.Index(context.Request.RequestURI, "/static") >= 0 || context.Request.RequestURI == "/" {
			fmt.Println()
		} else {
			cookie := context.Request.Header["Authorization"]
			timeout, username := services.CheckToken(cookie[0])
			fmt.Println(timeout, username)
			if username == "" {
				context.WriteString("无授权访问")
			}
		}
	})
}
