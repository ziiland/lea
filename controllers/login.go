package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/models"
	"lea/zllogs"
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
	models.InsertWorkerTableItem(account)

	account.WorkerId = "ZLD00001"
	account.Password = "888888"
	account.Name = "王晓光"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "经理"
	models.InsertWorkerTableItem(account)

	account.WorkerId = "ZLD00002"
	account.Password = "888888"
	account.Name = "殷骏"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "经理"
	models.InsertWorkerTableItem(account)	

	account.WorkerId = "ZLD00003"
	account.Password = "888888"
	account.Name = "张召"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "经理"
	models.InsertWorkerTableItem(account)

	account.WorkerId = "ZLD00004"
	account.Password = "888888"
	account.Name = "张三"
	account.Sex = "男"
	account.IdentifyNo = "123456789012345678"
	account.Title = "员工"
	models.InsertWorkerTableItem(account)	
	
}

/////////////////////////////////////////////////////////////////////////////////
func setSessionContent(c *LoginController) {
	ip := c.Ctx.Request.RemoteAddr
	id := c.GetString(ZLD_PARA_WORKERID)

	fmt.Println("login: ip=", ip)
	zllogs.WriteDebugLog("login: ip=%s, id=%s", ip, id)
	c.SetSession(ip, ZLD_STR_ON)
	c.SetSession(ZLD_PARA_WORKERID, id)
}

func (c *LoginController) Get() {
	// JUST FOR TEST
	//createWorkerTableItemForTest()

	// get the para
	workerId := c.GetString(ZLD_PARA_WORKERID)
	password := c.GetString(ZLD_PARA_PWD)

	item := new(LoginJsonData)
	// judgement the account
	if models.CheckWorkerLoginInfo(workerId, password) {
		// information correct!

		// create the session
		setSessionContent(c)

		// send back JSON data
		item.Errcode = 0;
		item.Page = "task_list.html"
	} else {
		item.Errcode = 1
		item.Page = ""
	}
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
