package routers

import (
	"logserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//首页
	beego.Router("/", &controllers.WelcomeController{})
	beego.Router("/join", &controllers.WelcomeController{}, "post:Join")
	//日志中心
	beego.Router("/logcenter", &controllers.LogCenterController{})
	//对外接口
	//更新设备信息
	beego.Router("/update_device", &controllers.UpdateCenterController{}, "post:UpdateDevice")
}
