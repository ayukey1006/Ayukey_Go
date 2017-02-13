package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type WelcomeController struct {
	beego.Controller
}

func (self *WelcomeController) Get() {
	self.TplName = "welcome.html"
}

func (self *WelcomeController) Join() {

	uname := self.Input().Get("uname")
	pwd := self.Input().Get("pwd")
	log.Println(uname)
	log.Println(pwd)

	self.Redirect("/center", 200)
}
