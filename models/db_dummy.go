package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ZldDummyData struct {
	Id		int32		`orm:"pk;auto"`
	Comment string
}

///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("register ZldDummyData table")
	orm.RegisterModel(new(ZldDummyData))
}