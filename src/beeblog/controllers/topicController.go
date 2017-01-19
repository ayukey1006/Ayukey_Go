package controllers

import (
	"beeblog/models"

	"strings"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (self *TopicController) Get() {
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	self.Data["IsTopic"] = true

	topics, err := models.GetAllTopics(false, "", "")
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
	tid := self.Input().Get("tid")

	title := self.Input().Get("title")
	content := self.Input().Get("content")
	category := self.Input().Get("category")
	tag := self.Input().Get("tag")

	if len(tid) == 0 {
		err = models.AddTopic(title, content, category, tag)
		if err != nil {
			beego.Error(err)
		} else {
			err = models.ModifyCategory(category, 1)
			if err != nil {
				beego.Error(err)
			}
		}
	} else {
		topic, err := models.GetTopic(tid)
		if err != nil {
			beego.Error(err)
		} else {
			err = models.ModifyTopic(tid, title, content, category, tag)
			if err != nil {
				beego.Error(err)
			} else {
				if topic.Category != category {
					err = models.ModifyCategory(topic.Category, -1)
					if err != nil {
						beego.Error(err)
					}

					err = models.ModifyCategory(category, 1)
					if err != nil {
						beego.Error(err)
					}
				}
			}
		}

	}

	self.Redirect("/topic", 302)
}

func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
}

func (self *TopicController) View() {
	self.TplName = "topic_view.html"

	id := self.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(id)
	if err != nil {
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = topic
	self.Data["Tid"] = id

	replies, err := models.GetAllReplies(id, true)
	if err != nil {
		beego.Error(err)
		return
	}

	self.Data["Replies"] = replies
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	self.Data["Tags"] = strings.Split(topic.Tags, " ")
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
