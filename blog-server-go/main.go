package main

import (
	_ "blog-server-go/routers"
	"github.com/astaxie/beego"
	"rachel-go/cache/redis"
	"rachel-go/common/filter"
)

func configuration() {
	// 跨域拦截器
	beego.InsertFilter("*", beego.BeforeRouter, filter.OriginFilter())
	// 请求头拦截
	beego.InsertFilter("*",beego.BeforeRouter, filter.CommonMiddlewareFilter())
	// 请求信息
	beego.InsertFilter("*",beego.BeforeRouter, filter.RequestParameter())
	//beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "redisServiceHost:6379"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379
}

func main() {
	configuration()
	redis.Init()
	beego.Run()
}
