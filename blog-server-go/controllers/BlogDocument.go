package controllers

import (
	"blog-server-go/models/config"
	blogEntity "blog-server-go/models/entity"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	constants2 "rachel-go/common/constants"
	"rachel-go/common/enums"
	"rachel-go/common/models/entity"
	"rachel-go/common/utils"
	"time"
)

type BlogDocumentController struct {
	beego.Controller
}

// Post 新增博客/*
func (document *BlogDocumentController) Post() {
	log.Println("请求开始：")
	startTime := time.Now()

	blogDocument :=&blogEntity.BlogDocument{}
	err := json.Unmarshal(document.Ctx.Input.RequestBody, blogDocument)
	if err != nil {
		log.Println("请求参数解析异常",err)
		_, _ = document.Ctx.ResponseWriter.Write(entity.ResultErrorMsg(constants2.ParameterException, constants2.ParameterExceptionMsg))
		return
	}

	blogDocument.CreateTime = utils.CurrentDate()
	blogDocument.UpdateTime = utils.CurrentDate()
	bytes, _ := json.Marshal(blogDocument)
	log.Println("新增博客",string(bytes))
	config.DB.Save(&blogDocument)

	log.Println("请求结束")
	log.Println("请求耗时：",time.Now().Sub(startTime))

	_, _ = document.Ctx.ResponseWriter.Write(entity.ResultSuccess())
}

// Put 更新博客/*
func (document *BlogDocumentController) Put() {
	log.Println("请求开始：")
	startTime := time.Now()

	blogDocument :=&blogEntity.BlogDocument{}
	err := json.Unmarshal(document.Ctx.Input.RequestBody, blogDocument)
	if err != nil {
		log.Println("请求参数解析异常",err)
		_, _ = document.Ctx.ResponseWriter.Write(entity.ResultErrorMsg(constants2.ParameterException, constants2.ParameterExceptionMsg))
		return
	}

	lastBlogDocument := &blogEntity.BlogDocument{}
	lastBlogDocument.ID = blogDocument.ID
	config.DB.Model(&lastBlogDocument).Find(&lastBlogDocument)
	blogDocument.UpdateTime = utils.CurrentDate()
	blogDocument.CreateTime = lastBlogDocument.CreateTime
	bytes, _ := json.Marshal(blogDocument)
	log.Println("更新博客",string(bytes))

	// 更新至数据库
	config.DB.Save(&blogDocument)

	log.Println("请求结束")
	log.Println("请求耗时：",time.Now().Sub(startTime))

	_, _ = document.Ctx.ResponseWriter.Write(entity.ResultSuccess())
}

// Delete 删除博客/*
func (document *BlogDocumentController) Delete() {
	log.Println("请求开始：")
	startTime := time.Now()

	documentId, _ := document.GetUint64(":documentId")
	blogDocument :=&blogEntity.BlogDocument{}
	if document == nil{
		log.Println("文档id为空",documentId)
		_, _ = document.Ctx.ResponseWriter.Write(entity.ResultErrorMsg(constants2.ParameterException, constants2.ParameterExceptionMsg))
		return
	}

	blogDocument.ID = documentId
	blogDocument.HasDel = enums.DataDelete
	blogDocument.UpdateTime = utils.CurrentDate()
	bytes, _ := json.Marshal(blogDocument)
	log.Println("删除博客",string(bytes))
	config.DB.Model(&blogDocument).Updates(map[string]interface{}{constants2.HasDel: blogDocument.HasDel, constants2.UpdateTime:blogDocument.UpdateTime})

	log.Println("请求结束")
	log.Println("请求耗时：",time.Now().Sub(startTime))

	_, _ = document.Ctx.ResponseWriter.Write(entity.ResultSuccess())
}
