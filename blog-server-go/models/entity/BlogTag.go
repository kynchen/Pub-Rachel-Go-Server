package entity

type BlogTag struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	ImgURL string `json:"img_url"`
	HasDel uint `json:"has_del"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
