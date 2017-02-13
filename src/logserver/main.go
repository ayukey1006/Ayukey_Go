package main

import (
	_ "logserver/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

