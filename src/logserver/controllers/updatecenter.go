package controllers

import (
	"encoding/json"
	"logserver/models"

	"github.com/astaxie/beego"
)

type UpdateCenterController struct {
	beego.Controller
}

func (self *UpdateCenterController) UpdateDevice() {
	beego.Info(self.Ctx.Request.RemoteAddr)

	var device models.Device
	json.Unmarshal(self.Ctx.Input.RequestBody, &device)
	beego.Info("传入参数:", device.Name)

	code := 0
	msg := "操作成功"

	if len(device.DeviceID) == 0 {
		code = -1
		msg = "device_id不能为空"
	}

	if len(device.Name) == 0 {
		code = -1
		msg = "device_name不能为空"
	}

	if len(device.IPAddress) == 0 {
		code = -1
		msg = "device_ip不能为空"
	}

	if code == 0 {
		err := models.UpdateDevice(device.DeviceID, device.Name, device.IPAddress)
		if err != nil {
			code = -1
			msg = err.Error()
		}
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

	self.Ctx.WriteString(string(b))

}

type LResponse struct {
	Code int
	Msg  string
}
