package main

import (
	"lea/database"
	_ "lea/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	o := orm.NewOrm()
	o.Using("default")
	database.CreateTable()

	beego.SetStaticPath("/mp", "static")
	beego.SetStaticPath("/view", "views")
	beego.Run()
}

