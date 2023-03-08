package models

type LoginResponse struct {
	Token   string `json:"token,omitempty"`
	Menus   []Memu `json:"menus,omitempty"`
	Routers string `json:"routers,omitempty"`
}

type Respond struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
}
