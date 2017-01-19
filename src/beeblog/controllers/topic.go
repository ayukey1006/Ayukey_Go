package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (self *TopicController) Get() {
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	self.Data["IsTopic"] = true

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		self.Data["Topics"] = topics
	}

	self.TplName = "topic.html"
}

func (self *TopicController) Post() {
	if !checkAccount(self.Ctx) {
		self.Redirect("/login", 302)
		return
	}

	var err error
	title := self.Input().Get("title")
	content := self.Input().Get("content")
	category := self.Input().Get("category")
	tid := self.Input().Get("tid")
	if len(tid) == 0 {
		err = models.AddTopic(title, content, category)
	} else {
		err = models.ModifyTopic(tid, title, content, category)
	}

	if err != nil {
		beego.Error(err)
	}

	self.Redirect("/topic", 302)
}

func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
}

func (self *TopicController) View() {
	id := self.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(id)
	if err != nil {
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = topic
	self.Data["Tid"] = id
	self.TplName = "topic_view.html"
}

func (self *TopicController) Modify() {
	tid := self.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = topic
	self.Data["Tid"] = tid
	self.TplName = "topic_modify.html"
}

func (self *TopicController) Delete() {
	if !checkAccount(self.Ctx) {
		self.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(self.Ctx.Input.Params()["0"])

	if err != nil {
		beego.Error(err)
	}

	self.Redirect("/", 302)
}
