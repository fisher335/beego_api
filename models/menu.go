package models

type Memu struct {
	Icon  string  `json:"icon,omitempty"`
	Index string  `json:"index,omitempty"`
	Title string  `json:"title,omitempty"`
	Subs  []*Memu `json:"subs,omitempty"`
}

type LoginResponse struct {
	Token   string `json:"token,omitempty"`
	Menus   []Memu `json:"menus,omitempty"`
	Routers string `json:"routers,omitempty"`
}
