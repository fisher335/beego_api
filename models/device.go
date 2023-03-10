package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"strings"
)

var (
	Objects map[string]*Device
)

type Device struct {
	Id     int    `orm:"pk,auto"`
	Name   string `orm:"size(256)"`
	No     string `orm:"size(256)"`
	Depart string `orm:"size(256)"`
	BaseModel
}

func init() {
	Objects = make(map[string]*Device)

}

func AddOne(device Device) (ObjectId string) {
	o := orm.NewOrm()
	o.Insert(device)
	return strconv.Itoa(device.Id)
}

func GetOne(DeviceId string) (*Device, error) {
	did, _ := strconv.Atoi(DeviceId)
	device := Device{Id: did}
	o := orm.NewOrm()
	err := o.Read(device)
	if err != nil {
		return nil, errors.New("id不存在")
	}

	return &device, nil
}

func GetAllDevice(page Pagination) ([]*Device, int) {
	o := orm.NewOrm()
	qs := o.QueryTable(&Device{}).Filter("name__icontains", strings.TrimSpace(page.Paras["name"]))
	total, _ := qs.Count()
	var devices []*Device
	_, err := qs.Limit(page.PageSize, (page.CurrentPage-1)*page.PageSize).All(&devices)
	if err != nil {
		return nil, 0
	}
	if err != nil {
		panic(errors.New("数据库查询错误"))
	}
	return devices, int(total)

}

func Update(decice *Device) (err error) {
	o := orm.NewOrm()
	o.Update(decice)
	return errors.New("ObjectId Not Exist")
}

func Delete(deviceid string) int {
	uid, _ := strconv.Atoi(deviceid)
	o := orm.NewOrm()
	device := Device{Id: uid}
	num, _ := o.Delete(device)
	return int(num)
}
