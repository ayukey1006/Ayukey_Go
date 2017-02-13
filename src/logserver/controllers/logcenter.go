package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type LogCenterController struct {
	beego.Controller
}

func (self *LogCenterController) Get() {
	log.Println("change page")
	self.TplName = "logcenter.html"
}
