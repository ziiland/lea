package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/models"
	"lea/utils/simplejson"
	"lea/zllogs"
	"strconv"
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
	item.Show = "true"
	models.InsertPriceTableItem(item)

	item.Name = "黄瓜种子"
	item.Kind = "G01_002"
	item.Price = 3.0	
	item.Show = "true"
	models.InsertPriceTableItem(item)

	item.Name = "翻地"
	item.Kind = "S02_001"
	item.Price = 4.0
	item.Show = "true"
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
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)

	item := new(PriceJsonData)
	item.Errcode = 1
	// para
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	fmt.Println("paraJSON=", paraJSON)

	price := models.NewZldPriceDBData()
	price.Name, _ = paraJSON.Get(common.ZLD_PARA_NAME).String()
	price.Kind, _ = paraJSON.Get(common.ZLD_PARA_KIND).String()
	price.Show, _ = paraJSON.Get(common.ZLD_PARA_SHOW).String()
	tPrice, _ := paraJSON.Get(common.ZLD_PARA_PRICE).String()
	price.Price, _ = strconv.ParseFloat(tPrice, 64)
	tDiscount, _ := paraJSON.Get(common.ZLD_PARA_DISCOUNT).String()
	price.Discount, _ = strconv.ParseFloat(tDiscount, 64)
	price.Comment, _ = paraJSON.Get(common.ZLD_PARA_COMMENT).String()

	if err := models.InsertPriceTableItem(price); err == nil {
		// Add log
		worker := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
		models.HandleStandardPriceLogItem(price.Kind, common.ZLD_CMD_ADD_RRICE, worker)
		item.Errcode = 0
	}
	c.Data["json"] = item
	c.ServeJSON()
}

func handleUpdatePriceCmd(c *PriceController) {
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)	

	item := new(PriceJsonData)
	item.Errcode = 1

	// para
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	fmt.Println("paraJSON=", paraJSON)
	price := models.NewZldPriceDBData()
	price.Name, _ = paraJSON.Get(common.ZLD_PARA_NAME).String()
	price.Kind, _ = paraJSON.Get(common.ZLD_PARA_KIND).String()
	price.Show, _ = paraJSON.Get(common.ZLD_PARA_SHOW).String()
	tPrice, _ := paraJSON.Get(common.ZLD_PARA_PRICE).String()
	price.Price, _ = strconv.ParseFloat(tPrice, 64)
	tDiscount, _ := paraJSON.Get(common.ZLD_PARA_DISCOUNT).String()
	price.Discount, _ = strconv.ParseFloat(tDiscount, 64)
	price.Comment, _ = paraJSON.Get(common.ZLD_PARA_COMMENT).String()

	if err := models.UpdatePriceItem(price); err == nil {
		// update log
		worker := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
		models.HandleStandardPriceLogItem(price.Kind, common.ZLD_CMD_UPDATE_PRICE, worker)	
		item.Errcode = 0		
	}	

	c.Data["json"] = item
	c.ServeJSON()	
}

func handleDeletePriceCmd(c *PriceController) {
	item := new(PriceJsonData)
	item.Errcode = 1

	kind := c.GetString(common.ZLD_PARA_KIND)
	if _, err := models.DeletePriceItem(kind); err == nil {
		// delete log
		worker := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
		models.HandleStandardPriceLogItem(kind, common.ZLD_CMD_DEL_PRICE, worker)			
		item.Errcode = 0
	}

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
	case common.ZLD_CMD_UPDATE_PRICE:
		handleUpdatePriceCmd(c)
	case common.ZLD_CMD_DEL_PRICE:
		handleDeletePriceCmd(c)
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
	case common.ZLD_CMD_ADD_RRICE:
		handleAddPriceCmd(c)		
	case common.ZLD_CMD_UPDATE_PRICE:
		handleUpdatePriceCmd(c)
	case common.ZLD_CMD_DEL_PRICE:
		handleDeletePriceCmd(c)		
	}
}
