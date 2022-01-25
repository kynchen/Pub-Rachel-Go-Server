package config

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

/*
 * 初始化连接
 */
func init()  {
	DB = createClient()
}
/*
 * 创建数据库连接
 */
func createClient() *gorm.DB {
	username := beego.AppConfig.String("mysql-user")
	password := beego.AppConfig.String("mysql-password")
	dbName := beego.AppConfig.String("mysql-dbname")
	host := beego.AppConfig.String("mysql-host")
	maxIdleConn, _ := beego.AppConfig.Int("mysql-max-idle-conn")
	maxOpenConn, _ := beego.AppConfig.Int("mysql-max-open-conn")

	args := username+":"+password+"@tcp("+host+")/"+dbName+"?charset=utf8"
	log.Println("mysql连接信息",args)
	open, err := gorm.Open("mysql", args)
	if err !=nil {
		log.Println("初始化mysql连接失败")
		panic(err)
	}
	// 空闲连接池连接数
	open.DB().SetMaxIdleConns(maxIdleConn)
	// 最大连接数
	open.DB().SetMaxOpenConns(maxOpenConn)

	return open
}