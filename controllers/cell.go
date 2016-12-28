package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/models"
	//"lea/utils/simplejson"
	"lea/zllogs"
	//"strconv"
	//"strings"
	//"time"
)


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type CellController struct {
	beego.Controller
}

type CellJsonData struct {
	//Prices  			*[]models.ZldPriceData
	Errcode             int
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func handleBindNFCCmd(c *CellController) {
	farm := c.GetString(common.ZLD_PARA_FARM)
	cell := c.GetString(common.ZLD_PARA_CELL)
	nfc := c.GetString(common.ZLD_PARA_NFC)

	item := new(CellJsonData)
	item.Errcode = 1
	// write to db
	if models.BindCellNFCTableItem(farm, cell, nfc) == 0 {
		item.Errcode = 0;
	}

	c.Data["json"] = item
	c.ServeJSON()		
}

func (c *CellController) Get() {
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("Cell command=", command)
	zllogs.WriteDebugLog("GET request of cell page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_CELL_BINDNFC:
		handleBindNFCCmd(c)
	}
}

func (c *CellController) Post() {
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("Cell command=", command)
	zllogs.WriteDebugLog("POST request of cell page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_CELL_BINDNFC:
		handleBindNFCCmd(c)	
	}
}
