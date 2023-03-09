package services

import (
	"beego_api/models"
	"crypto/md5"
	"errors"
	"fmt"
)

func AuthenticateUserForLogin(loginName, password string) (models.User, error) {
	user := models.User{}
	if len(password) == 0 || len(loginName) == 0 {
		return user, errors.New("Error:用户或者密码为空")
	}
	data := []byte(password)
	has := md5.Sum(data)
	password = fmt.Sprintf("%x", has)
	v, err := models.GetUserByLoginName(loginName) //数据库查询语句。自己写的
	if err != nil {
		return user, errors.New("Error:未找到该用户")

	} else if v.Password != password {
		return user, errors.New("Error:密码错误")

	} else {
		return v, nil
	}
}

func GetUserList(page models.Pagination) map[string]interface{} {
	users, total := models.GetAllUsers(page)
	var result = make(map[string]interface{})
	result["list"] = users
	result["currentPage"] = page.CurrentPage
	result["pageSize"] = page.PageSize
	result["total"] = total
	return result
}
