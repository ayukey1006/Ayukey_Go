package controllers

import (
	"beeblog/models"

	"time"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (self *ReplyController) Add() {
	tid := self.Input().Get("tid")
	nickname := self.Input().Get("nickname")
	content := self.Input().Get("content")

	reply, err := models.AddReply(tid, nickname, content)
	if err != nil {
		beego.Error(err)
	} else {
		models.ModifyTopicReply(tid, true, reply.Created)
	}

	self.Redirect("/topic/view/"+tid, 301)
}

func (self *ReplyController) Delete() {
	rid := self.Input().Get("rid")
	tid := self.Input().Get("tid")

	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)
	} else {
		models.ModifyTopicReply(tid, false, time.Now())
	}

	self.Redirect("/topic/view/"+tid, 302)
}
