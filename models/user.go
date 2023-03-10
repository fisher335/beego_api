package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

var (
	UserList map[string]*User
)

type BaseModel struct {
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);null"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);null"`
}

type User struct {
	Id       int64  `orm:"auto;pk"  json:"id,omitempty"`
	Username string `orm:"unique;size(255)" json:"login_name,omitempty" `
	Name     string `orm:"size(255);null" json:"name,omitempty"`
	Type     string `orm:"size(255);default(普通)" json:"type,omitempty"`
	Email    string `orm:"size(255);null" json:"email,omitempty"`
	Phone    string `orm:"size(255);null" json:"phone,omitempty"`
	Password string `orm:"size(255)" json:"password,omitempty"`
	Test     int    `orm:"default(1)"`
	BaseModel
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

func GetAllUsers(page Pagination) ([]*User, int) {
	o := orm.NewOrm()
	qs := o.QueryTable(&User{}).Filter("username__icontains", strings.TrimSpace(page.Paras["name"]))
	total, _ := qs.Count()
	var users []*User
	_, err := qs.Limit(page.PageSize, (page.CurrentPage-1)*page.PageSize).All(&users)
	if err != nil {
		return nil, 0
	}
	if err != nil {
		panic(errors.New("数据库查询错误"))
	}
	return users, int(total)
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

func DeleteUser(uid string) {
	delete(UserList, uid)
}
func GetUserByLoginName(name string) (User, error) {
	user := User{Username: name}
	o := orm.NewOrm()
	err := o.Read(&user, "username")
	if err != nil {
		panic(err)
	}
	return user, nil
}
