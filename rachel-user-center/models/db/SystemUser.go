package db

type SystemUser struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Sex uint `json:"sex"`
	Sign string `json:"sign"`
	Avatar string `json:"avatar"`
	Nation string `json:"nation"`
	Province string `json:"province"`
	City string `json:"city"`
	Region string `json:"region"`
	HasDel uint `json:"has_del"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}