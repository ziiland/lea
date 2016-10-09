package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/models"
)

func init() {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	dbname := beego.AppConfig.String("dbname")
	//host := beego.AppConfig.String("host")
	//port := beego.AppConfig.String("port")

	fmt.Println("package database: init function")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", 
		fmt.Sprintf("%s:%s@/%s?charset=utf8", username, password, dbname), 30)
	//orm.RegisterDataBase("default", "mysql", 
	//	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbname), 30)
}

func CreateTable() {
	name := "default"
	force := false
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error(err)
	}

	// create table by SQL
	models.CreateOrderContentTable()
}






