package controllers

import (
	"github.com/astaxie/beego"
	"lea/models"
)


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type LoginController struct {
	beego.Controller
}

type LoginJsonData struct {
	Page 				string
	//Title               string
	Errcode             int
}

///////////////////////////////////////////////////////////////////////////////
func createWorkerTableItemForTest() {
	// create Admin account
	account := models.NewZldWorkerDBData()

	account.WorkerId = "Admin"
	account.Password = "Admin"
	account.Name = "管理员"
	account.Title = "管理员"
	account.Comment = "管理员"
	models.DoInsertZldWorkerTableItem(account)

	account.WorkerId = "ZLD00001"
	account.Password = "888888"
	account.Name = "王晓光"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "经理"
	models.DoInsertZldWorkerTableItem(account)

	account.WorkerId = "ZLD00002"
	account.Password = "888888"
	account.Name = "殷骏"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "经理"
	models.DoInsertZldWorkerTableItem(account)	

	account.WorkerId = "ZLD00003"
	account.Password = "888888"
	account.Name = "张召"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "经理"
	models.DoInsertZldWorkerTableItem(account)

	account.WorkerId = "ZLD00004"
	account.Password = "888888"
	account.Name = "张三"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "员工"
	models.DoInsertZldWorkerTableItem(account)	
	
}

func (c *LoginController) Get() {
	// JUST FOR TEST
	//createWorkerTableItemForTest()

	// get the para
	workerId := c.GetString(ZLD_PARA_WORKERID)
	password := c.GetString(ZLD_PARA_PWD)

	item := new(LoginJsonData)
	item.Errcode = 1;
	// judgement the account
	if models.CheckWorkerLoginInfo(workerId, password) {
		// information correct!
		item.Errcode = 0;
	}

	item.Page = ""
	c.Data["json"] = item
	c.ServeJSON()	
}

func (c *LoginController) Post() {
	// get the para
	workerId := c.GetString(ZLD_PARA_WORKERID)
	password := c.GetString(ZLD_PARA_PWD)

	item := new(LoginJsonData)
	item.Errcode = 1;
	// judgement the account
	if models.CheckWorkerLoginInfo(workerId, password) {
		// information correct!
		item.Errcode = 0;
	}
		
	item.Page = ""
	c.Data["json"] = item
	c.ServeJSON()
}
