package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

var (
	UserList map[string]*User
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

type User struct {
	Id       int    `orm:"id"`
	Username string `orm:"username"`
	Password string `orm:"password"`
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func AddUser(u User) string {
	o := orm.NewOrm()
	id, _ := o.Insert(&u)
	return strconv.FormatInt(id, 10)
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}

		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM user WHERE username = ?", username).QueryRow(&user)
	if err != nil {
		panic(err)
	}
	if user.Password == password {
		return true
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
