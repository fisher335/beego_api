package common

type respond struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
}

func OK(data interface{}) respond {
	r := respond{}
	r.Code = 200
	r.Data = data
	r.Status = "success"
	r.Message = ""
	return r
}

type LoginResponse struct {
	Token   string `json:"token,omitempty"`
	Menus   string `json:"menus,omitempty"`
	Routers string `json:"routers,omitempty"`
}


