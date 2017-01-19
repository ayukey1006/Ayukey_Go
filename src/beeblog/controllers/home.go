package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.Data["IsHome"] = true
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	self.Data["Topics"] = topics
	self.TplName = "home.html"
}
