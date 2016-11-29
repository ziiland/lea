package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/zllogs"
)

type MainController struct {
	beego.Controller
}

type CommonCmdJsonData struct {
	Login  		string 
	Worker  	string   // the worker who login 
	Title		string   // the title
	FarmId		string   // farm which worker belong

	ACK         string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func handleLoadParaCmd(c *beego.Controller) {
	// send back json including Login/Order/Node
	fmt.Println("handleLoadParaCmd")
	item := new(CommonCmdJsonData)

	sess := c.StartSession()
	fmt.Printf("Session: id=%v, worker=%s\n", sess.SessionID(), sess.Get(common.ZLD_PARA_WORKER))

	ip := c.Ctx.Request.RemoteAddr
	fmt.Println("ip=", ip)
	if c.GetSession(ip) == common.ZLD_STR_ON {
		item.Login = common.ZLD_STR_ON
		item.Worker = (c.GetSession(common.ZLD_PARA_WORKER)).(string)
		item.Title = (c.GetSession(common.ZLD_PARA_TITLE)).(string)
		fmt.Printf("%s is LOGIN!\n", item.Worker)
		zllogs.WriteDebugLog("%s is LOGIN!", item.Worker)
	} else {
		//item.Login = 
		fmt.Println("worker not login!")
		zllogs.WriteErrorLog("worker not login!")
	}

	item.ACK = common.ZLD_STR_OK	
	c.Data["json"] = item
	c.ServeJSON()
}

func handleUnloadCmd(c *beego.Controller) {
	// record the logout action
	fmt.Println("handleUnloadCmd")
	// dbNode := models.NewNodeInfoDBData()

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("r=", r)
	// 		dglogs.WriteErrorLog("handleNodeUnloadCmd: runtime error catch %v", r)
	// 	}
	// 	item := new(NodeCmdJsonData)
	// 	item.ACK = DG_PARA_OK
	// 	c.Data["json"] = item
	// 	c.ServeJSON()		
	// }()		

	// dbNode.Ip = c.Ctx.Request.RemoteAddr
	// dbNode.Node = (c.GetSession(DG_PARA_NODE)).(string)
	// dbNode.Action = DG_STR_LOGOUT

	// order := (c.GetSession(DG_PARA_ORDER)).(string)
	// models.InsertNodeInfoTableItem(order, dbNode)
	// dglogs.WriteDebugLog("%s is LOGOUT!", dbNode.Node)

	// del session
	ip := c.Ctx.Request.RemoteAddr
	c.DelSession(ip)

	item := new(CommonCmdJsonData)
	item.ACK = common.ZLD_STR_OK
	c.Data["json"] = item
	c.ServeJSON()
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
