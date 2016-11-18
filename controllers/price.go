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
type PriceController struct {
	beego.Controller
}

type PriceJsonData struct {
	Prices  			*[]models.ZldPriceData
	Errcode             int
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func createPriceTableItemForTest() {
	// create some price items
	item := models.NewZldPriceDBData()
	item.Name = "青菜种子"
	item.Kind = "G01_001"
	item.Price = 2.0
	item.Discount = 1
	models.InsertPriceTableItem(item)

	item.Name = "黄瓜种子"
	item.Kind = "G01_002"
	item.Price = 3.0	
	models.InsertPriceTableItem(item)

	item.Name = "翻地"
	item.Kind = "S02_001"
	item.Price = 4.0
	models.InsertPriceTableItem(item)
}

func loadAllPriceInfo(c *PriceController) {
	item := new(PriceJsonData)
	item.Errcode = 1
	if prices, err := models.QueryAllPriceTableItem(); err == nil {
		slice := make([]models.ZldPriceData, 0)
		for _, v := range prices {
			slice = append(slice, v)
		}
		item.Prices = &slice
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()	
}

func handleLoadPriceCmd(c *PriceController) {
	title := (c.GetSession(common.ZLD_PARA_TITLE)).(string)

	if title == common.ZLD_STR_ADMIN {
		loadAllPriceInfo(c)
	} else {
		item := new(PriceJsonData)
		item.Errcode = 1
		c.Data["json"] = item
		c.ServeJSON()			
	}
}

func handleAddPriceCmd(c *PriceController) {
	item := new(PriceJsonData)
	item.Errcode = 1

	c.Data["json"] = item
	c.ServeJSON()
}

func (c *PriceController) Get() {
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("price command=", command)
	zllogs.WriteDebugLog("GET request of price page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_LOAD_PRICE:
		handleLoadPriceCmd(c)
	case common.ZLD_CMD_ADD_RRICE:
		handleAddPriceCmd(c)
	}
}

func (c *PriceController) Post() {
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("price command=", command)
	zllogs.WriteDebugLog("POST request of price page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_LOAD_PRICE:
		handleLoadPriceCmd(c)
	}
}
