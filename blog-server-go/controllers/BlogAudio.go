package controllers

import (
	"blog-server-go/models/config"
	blogEntity "blog-server-go/models/entity"
	"github.com/astaxie/beego"
	"log"
	"rachel-go/common/enums"
	"rachel-go/common/models/entity"
	"time"
)

type BlogAudioController struct {
	beego.Controller
}

// Get 获取歌单/*
func (document *BlogAudioController) Get() {
	log.Println("请求开始：")
	startTime := time.Now()

	addData, _ := document.GetBool("allData")
	var blogAudioList []blogEntity.BlogAudio
	if addData {
		// 获取所有歌单
		config.DB.Find(&blogAudioList)
	}else{
		// 获取未被删除的所有歌单
		config.DB.Where("has_del = ?", enums.DataNormal).Find(&blogAudioList)
	}

	log.Println("请求结束")
	log.Println("请求耗时：",time.Now().Sub(startTime))

	_, _ = document.Ctx.ResponseWriter.Write(entity.ResultSuccessData(blogAudioList))
}