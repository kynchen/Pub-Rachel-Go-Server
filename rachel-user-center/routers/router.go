package routers

import (
	"github.com/astaxie/beego"
	"rachel-user-center/controllers"
)

func init() {
    beego.Router("/system/user/:id", &controllers.UserInfoController{})
}
