package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	// 需要在init中注册定义的model
	//orm.RegisterDriver("sqlite", orm.DRSqlite)
	//orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wish@tcp(172.16.120.198:3306)/cect?charset=utf8&loc=Local")

	orm.RegisterModel(new(User), new(Device))
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}
