package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/models"
	"lea/utils/simplejson"
	"lea/zllogs"
	//"strings"
	//"strconv"
	//"time"
)


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type PacketController struct {
	beego.Controller
}

type PacketJsonData struct {
	Sender 				string
	Receiver            string
	RMobile				string
	Address				string
	ExpressNo			string
	Errcode             int
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// func queryMatchItemNums(workerId, title string) (int64, error){
// 	if title == common.ZLD_STR_WORKER {
// 		// only query task NOT CLOSED & worker/checker == workerId
// 		return models.QueryWorkerItemNums(workerId)
// 	} else {
// 		// query task NOT CLOSED
// 		return models.QueryAllOpenItemNums()
// 	}
// }
func createPacketTableItemForTest() {
	packet := models.NewZldPacketDBData()
	packet.TaskId = "SHA01A000A0001N14834315850000000"
	packet.Sender = "殷骏"
	packet.Receiver = "williamzhang"
	packet.Address = "上海市莲花路1978号E栋503"
	packet.RMobile = "12809876543"
	models.InsertPacketTableItem(packet)
}

func handleAddPacketCmd(c *PacketController) {
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)

	// para
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	fmt.Println("paraJSON=", paraJSON)
	models.AddPacketJSONString(paraJSON)

	item := new(PacketJsonData)
	item.Errcode = 0
	c.Data["json"] = item
	c.ServeJSON()	
}

func handleGetPacketCmd(c *PacketController) {
	taskId := c.GetString(common.ZLD_PARA_TASKID)
	expressNo := c.GetString(common.ZLD_PARA_EXPNO)
	packetId, _ := c.GetInt(common.ZLD_PARA_PACKEDID)
	fmt.Printf("handleGetPacketCmd: TaskId=%s, ExpressNo=%s, PacketId=%s\n", taskId, expressNo, packetId)

	item := new(PacketJsonData)
	item.Errcode = 1

	packet := models.NewZldPacketDBData()
	packet.TaskId = taskId
	packet.ExpressNo = expressNo
	packet.Id = int32(packetId)
	if err := models.SelectPacketTableItemWithTaskId(packet); err == nil {
		item.Sender = packet.Sender
		item.Receiver = packet.Receiver
		item.Address = packet.Address
		item.RMobile = packet.RMobile
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()	
}

func handleSetPacketCmd(c *PacketController) {
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)

	// para
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	fmt.Println("paraJSON=", paraJSON)

	item := new(PacketJsonData)
	item.Errcode = 1
	if err := models.UpdatePacketJSONString(paraJSON); err == nil {
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()	
}

func handleSetExpressNoCmd(c *PacketController) {
	taskId := c.GetString("TaskId")
	expNo := c.GetString("ExpressNo")
	fmt.Printf("handleSetExpressNoCmd:taskId=%s, expNo=%s\n", taskId, expNo)	

	item := new(PacketJsonData)
	item.Errcode = 1	
	if err := models.UpdatePacketExpressNo(taskId, expNo); err == nil {
		item.Errcode = 0
	}
	c.Data["json"] = item 
	c.ServeJSON()

}

///////////////////////////////////////////////////////////////////////////////
func (c *PacketController) Get() {
	// get the para
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("task command=", command)
	zllogs.WriteDebugLog("GET request of packet page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_ADD_PACKET:
		handleAddPacketCmd(c)
	case common.ZLD_CMD_GET_PACKET:
		handleGetPacketCmd(c)
	case common.ZLD_CMD_SET_PACKET:
		handleSetPacketCmd(c)
	case common.ZLD_CMD_SET_EXPNO:
		handleSetExpressNoCmd(c)
	}	
}

func (c *PacketController) Post() {
	// get the para
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("task command=", command)
	zllogs.WriteDebugLog("POST request of packet page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_ADD_PACKET:
		handleAddPacketCmd(c)		
	case common.ZLD_CMD_GET_PACKET:
		handleGetPacketCmd(c)
	case common.ZLD_CMD_SET_PACKET:
		handleSetPacketCmd(c)
	case common.ZLD_CMD_SET_EXPNO:
		handleSetExpressNoCmd(c)
	}	
}
