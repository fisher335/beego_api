package services

import "beego_api/models"

func GetAllDevices(page models.Pagination) map[string]interface{} {
	devices, total := models.GetAllDevice(page)
	var result = make(map[string]interface{})
	result["list"] = devices
	result["currentPage"] = page.CurrentPage
	result["pageSize"] = page.PageSize
	result["total"] = total
	return result

}
