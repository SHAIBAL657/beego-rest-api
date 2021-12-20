package main

import (
	"sara/db"
	_ "sara/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	db.InitDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
