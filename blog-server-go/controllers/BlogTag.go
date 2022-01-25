package controllers

import (
	"blog-server-go/models/config"
	blogEntity "blog-server-go/models/entity"
	"github.com/astaxie/beego"
	"log"
	"rachel-go/common/models/entity"
	"time"
)

type BlogTagController struct {
	beego.Controller
}

// Get 获取标签/*
func (document *BlogTagController) Get() {
	log.Println("请求开始：")
	startTime := time.Now()

	var blogTagList []blogEntity.BlogTag
	tagName := document.GetString("name")
	config.DB.Where("name LIKE ?", "%"+tagName+"%").Find(&blogTagList)

	log.Println("请求结束")
	log.Println("请求耗时：",time.Now().Sub(startTime))

	_, _ = document.Ctx.ResponseWriter.Write(entity.ResultSuccessData(blogTagList))
}