package controllers

import (
	"beego_api/common"
	"beego_api/models"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// UserController
// Operations about Users
type UserController struct {
	beego.Controller
}

// Post
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// GetAll @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
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
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		l := models.LoginResponse{
			Token:   "123123123",
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

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
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
