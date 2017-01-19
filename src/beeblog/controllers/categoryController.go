package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (self *CategoryController) Get() {

	op := self.Input().Get("op")
	switch op {
	case "add":
		name := self.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCateGory(name)
		if err != nil {
			beego.Error(err)
		}
		self.Redirect("/category", 301)
		return
	case "del":
		id := self.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		self.Redirect("/category", 301)
		return
	}

	self.Data["IsCategory"] = true

	var err error
	self.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	self.TplName = "category.html"
}
