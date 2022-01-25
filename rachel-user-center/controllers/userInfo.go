package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"rachel-go/common/models/entity"
	"rachel-user-center/models/config"
	"rachel-user-center/models/db"
	"time"
)

type UserInfoController struct {
	beego.Controller
}

func (document *UserInfoController) Get() {

	log.Println("Start --> 获取用户信息")
	startTime := time.Now()

	userId,_ := document.GetUint64(":id")
	user := db.SystemUser{
		ID: userId,
	}
	config.DB.Where(userId).First(&user)

	log.Println("End --> 获取用户信息")
	log.Println("耗时 --> ",time.Now().Sub(startTime))

	document.Ctx.ResponseWriter.Write(entity.ResultSuccessData(user))
}
