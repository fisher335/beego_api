package controllers

import (
	"beego_api/common"
	"beego_api/models"
	"beego_api/services"
	"crypto/md5"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	data := []byte(user.Password)
	has := md5.Sum(data)
	user.Password = fmt.Sprintf("%x", has)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

func (u *UserController) GetAll() {
	var page = models.Pagination{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &page)
	if err != nil {
		panic(err)
	}
	var data = services.GetUserList(page)
	u.Data["json"] = common.OK(data)
	u.ServeJSON()
}

func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

func (u *UserController) Delete() {
	uid := u.GetString("uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	user, _ := services.AuthenticateUserForLogin(username, password)
	if user.Username != "" {
		l := models.LoginResponse{
			Token:   services.CreateToken(username),
			Menus:   u.GetMenu(),
			Routers: "/home_/users_/user/info_/test",
		}

		re := common.OK(l)
		result, _ := json.Marshal(re)
		fmt.Print(string(result))
		u.Data["json"] = &re
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

func (u *UserController) GetMenu() []models.Memu {

	var str = []byte(`[{"icon":"el-icon-setting","index":"/home","title":"首页"},{"icon":"el-icon-edit","index":"2","title":"用户管理","subs":[{"icon":"el-icon-edit","index":"/users","title":"用户列表"}]},{"icon":"el-icon-setting","index":"/scaner","title":"主机扫描"},{"icon":"el-icon-zoom-in","index":"/industrial","title":"漏洞检查"},{"icon":"el-icon-setting","index":"/shell","title":"webShell"},{"icon":"el-icon-share","index":"/vis","title":"拓扑结构"},{"icon":"el-icon-delete","index":"/luyou","title":"设备管理"},{"icon":"el-icon-bell","index":"/test","title":"测试管理"}]`)
	var menus []models.Memu
	err := json.Unmarshal(str, &menus)
	if err != nil {
		panic(err)
	}
	return menus
}
