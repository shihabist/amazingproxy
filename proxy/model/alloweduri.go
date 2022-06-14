package model

type AllowedListData struct {
	AllowedList []AllowedUri `json:"allowedList"`
	Rules       Rules        `json:"rules,omitempty"`
}
type AllowedUri struct {
	ID     string `json:"id,omitempty"`
	Uri    string `json:"uri"`
	Method string `json:"method"`
}
type Rules struct {
	ID  string `json:"id,omitempty"`
	Num string `json:"num,omitempty"`
}
