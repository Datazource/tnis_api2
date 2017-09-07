package main

import (
	_ "tnis_api2/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver(beego.AppConfig.String("sqldriver"), orm.DRMySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("sqldriver"), beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/"+beego.AppConfig.String("mysqldb")+"?charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
