package controllers

import (
	"beego_api/common"
	"beego_api/models"
	"beego_api/services"
	"encoding/json"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type DeviceController struct {
	beego.Controller
}

func (o *DeviceController) Post() {
	var ob models.Device
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

func (o *DeviceController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

func (o *DeviceController) GetAll() {
	var page = models.Pagination{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &page)
	if err != nil {
		panic(err)
	}
	data := services.GetAllDevices(page)
	o.Data["json"] = common.OK(data)
	o.ServeJSON()
}

func (o *DeviceController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Device
	ob.Id, _ = strconv.Atoi(objectId)
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err := models.Update(&ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

func (o *DeviceController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
