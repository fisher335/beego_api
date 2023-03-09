package common

type Respond struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
}

func OK(data interface{}) Respond {
	r := Respond{}
	r.Code = 200
	r.Data = data
	r.Status = "success"
	r.Message = ""
	return r
}
