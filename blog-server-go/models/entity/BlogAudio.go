package entity

type BlogAudio struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Cover string `json:"cover"`
	Author string `json:"author"`
	AudioURL string `json:"audio_url"`
	AudioLrcURL string `json:"audio_lrc_url"`
	AudioType uint `json:"audio_type"`
	HasDel uint `json:"has_del"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
