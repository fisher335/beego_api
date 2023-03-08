package common

import "beego_api/models"

func OK(data interface{}) models.Respond {
	r := models.Respond{}
	r.Code = 200
	r.Data = data
	r.Status = "success"
	r.Message = ""
	return r
}
