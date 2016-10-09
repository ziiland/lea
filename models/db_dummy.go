package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
)

type DgDummyData struct {
	Id		int32		`orm:"pk;auto"`
	Comment string
}

///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("register DgDummyData table")
	orm.RegisterModel(new(DgDummyData))
}