package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/zllogs"
)

type CommonCmdJsonData struct {
	Login  		string 
	WorkerId	string

	ACK         string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func handleLoadParaCmd(c *beego.Controller) {
	// send back json including Login/Order/Node
	fmt.Println("handleLoadParaCmd")
	item := new(CommonCmdJsonData)

	ip := c.Ctx.Request.RemoteAddr
	fmt.Println("ip=", ip)
	if c.GetSession(ip) == ZLD_STR_ON {
		item.Login = ZLD_STR_ON
		item.WorkerId = (c.GetSession(ZLD_PARA_WORKERID)).(string)
		fmt.Printf("%s is LOGIN!\n", item.WorkerId)
		zllogs.WriteDebugLog("%s is LOGIN!", item.WorkerId)
	} else {
		//item.Login = 
		fmt.Println("worker not login!")
		zllogs.WriteErrorLog("worker not login!")
	}

	item.ACK = ZLD_STR_OK	
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
	item.ACK = ZLD_STR_OK
	c.Data["json"] = item
	c.ServeJSON()
}