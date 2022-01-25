package entity

type BlogDocument struct {
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Cover string `json:"cover"`
	HasDel uint `json:"has_del"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
