package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.TplName = "home.html"

	self.Data["IsHome"] = true
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	cate := self.Input().Get("cate")
	tag := self.Input().Get("tag")
	topics, err := models.GetAllTopics(true, cate, tag)
	if err != nil {
		beego.Error(err)
	}
	self.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	self.Data["Categories"] = categories
}
