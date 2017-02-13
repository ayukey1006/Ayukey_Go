package controllers

import (
	"encoding/json"
	"logserver/models"

	"log"

	"github.com/astaxie/beego"
)

type UpdateCenterController struct {
	beego.Controller
}

func (self *UpdateCenterController) UpdateDevice() {
	id := self.Input().Get("device_id")
	name := self.Input().Get("device_name")
	ip := self.Input().Get("device_ip")

	code := 0
	msg := "操作成功"

	if len(id) == 0 {
		code = -1
		msg = "device_id不能为空"
	}

	if len(name) == 0 {
		code = -1
		msg = "device_name不能为空"
	}

	if len(ip) == 0 {
		code = -1
		msg = "device_ip不能为空"
	}

	err := models.UpdateDevice(id, name, ip)
	if err != nil {
		code = -1
		msg = err.Error()
	}

	rs := LResponse{
		Code: code,
		Msg:  msg,
	}

	b, err := json.Marshal(rs)
	self.Ctx.ResponseWriter.WriteHeader(200)
	if err != nil {
		self.Ctx.WriteString(err.Error())
		return
	}

	log.Println(string(b))
	self.Ctx.WriteString(string(b))

}

type LResponse struct {
	Code int    
	Msg  string 
}
