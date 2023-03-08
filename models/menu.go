package models

type Memu struct {
	Icon  string  `json:"icon,omitempty"`
	Index string  `json:"index,omitempty"`
	Title string  `json:"title,omitempty"`
	Subs  []*Memu `json:"subs,omitempty"`
}
