package main

import (
	_ "logserver/routers"

	"github.com/astaxie/beego"

	"logserver/models"

	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
