package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
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


///////////////////////////////////////////////////////////////////////////////
func setSessionContent(c *LoginController) {
	//ip := c.Ctx.Request.RemoteAddr
	id := c.GetString(common.ZLD_PARA_WORKER)

	//fmt.Println("login: ip=", ip)
	//zllogs.WriteDebugLog("login: ip=%s, id=%s", ip, id)
	sess := c.StartSession()
	sessId := sess.SessionID()
	fmt.Printf("Session: id=%v\n", sess.SessionID())
	zllogs.WriteDebugLog("login: session=%s, worker=%s", sessId, id)
	c.SetSession(sessId, common.ZLD_STR_ON)
	c.SetSession(common.ZLD_PARA_WORKER, id)

	worker := models.NewZldWorkerDBData()
	worker.WorkerId = id
	if err := models.SelectWorkerTableItem(worker); err == nil {
		fmt.Println("worker.Title=", worker.Title)
		c.SetSession(common.ZLD_PARA_TITLE, worker.Title)
		//c.SetSession(ZLD_PARA_FARM, worker.Farm)
	}
}

func handleLoginCmd(c *LoginController) {
	workerId := c.GetString(common.ZLD_PARA_WORKER)
	password := c.GetString(common.ZLD_PARA_PWD)

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

func (c *LoginController) Get() {
	// JUST FOR TEST
	//createWorkerTableItemForTest()
	//createTaskTableItemForTest()
	//createTaskLogTableItemForTest()
	//createPriceTableItemForTest()
	//createPacketTableItemForTest()
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("login command=", command)
	zllogs.WriteDebugLog("GET request of login page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_LOGIN:
		handleLoginCmd(c)
	}
}

func (c *LoginController) Post() {
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("login command=", command)
	zllogs.WriteDebugLog("POST request of login page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_LOGIN:
		handleLoginCmd(c)
	}
}
