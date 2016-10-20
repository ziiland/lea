package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/models"
	"lea/zllogs"
	//"strconv"
	//"time"
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


/////////////////////////////////////////////////////////////////////////////////
func setSessionContent(c *LoginController) {
	ip := c.Ctx.Request.RemoteAddr
	id := c.GetString(ZLD_PARA_WORKER)

	fmt.Println("login: ip=", ip)
	zllogs.WriteDebugLog("login: ip=%s, id=%s", ip, id)
	c.SetSession(ip, ZLD_STR_ON)
	c.SetSession(ZLD_PARA_WORKER, id)

	worker := models.NewZldWorkerDBData()
	worker.WorkerId = id
	if err := models.SelectWorkerTableItem(worker); err == nil {
		fmt.Println("worker.Title=", worker.Title)
		c.SetSession(ZLD_PARA_TITLE, worker.Title)
		//c.SetSession(ZLD_PARA_FARM, worker.Farm)
	}
}

func (c *LoginController) Get() {
	// JUST FOR TEST
	//createWorkerTableItemForTest()
	//createTaskTableItemForTest()

	// get the para
	workerId := c.GetString(ZLD_PARA_WORKER)
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
	workerId := c.GetString(ZLD_PARA_WORKER)
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
