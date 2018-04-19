package main

import (
	_ "github.com/CJianWei/PPGo_ApiAdmin/models"
	_ "github.com/CJianWei/PPGo_ApiAdmin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
