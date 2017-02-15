package controllers

import (
	"log"
	"logserver/models"

	"github.com/astaxie/beego"
)

type LogCenterController struct {
	beego.Controller
}

func (self *LogCenterController) Get() {
	log.Println("change page")
	self.TplName = "logcenter.html"

	devices, err := models.GetAllDevice()
	if err != nil {
		beego.Error(err)
	}

	if len(devices) > 0 {
		self.Data["HasDevice"] = true
	} else {
		self.Data["HasDevice"] = false
	}

	self.Data["Devices"] = devices

	beego.Info(self.Data["Devices"])
	for _, device := range devices {
		beego.Info(device.Name)
	}

}
