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
type WorkerController struct {
	beego.Controller
}

type WorkerJsonData struct {
	Tasks 				*[]models.ZldTaskData
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

///////////////////////////////////////////////////////////////////////////////
func handleWorkerLoadTaskCmd(c *WorkerController) {
	workerId := (c.GetSession(ZLD_PARA_WORKER)).(string)
	title := (c.GetSession(ZLD_PARA_TITLE)).(string)

	fmt.Println("handleWorkerLoadTaskCmd: worker=", workerId)
	if title == "员工" {
		fmt.Println("title=", title)
	} else if title == "经理" {
		fmt.Println("title=", title)
	}

	item := new(WorkerJsonData)
	//slice := make([]models.ZldTaskData, 1)
	if num, err := models.QueryMatchItemNums(workerId, "SHA001", title); err == nil {
		slice := make([]models.ZldTaskData, num)
		item.Tasks = &slice
		models.SelectTaskTableItemsWithFarmId(workerId, "SHA001", title, item.Tasks)
	}	
	item.Errcode = 0;

	c.Data["json"] = item
	c.ServeJSON()
}

///////////////////////////////////////////////////////////////////////////////
func (c *WorkerController) Get() {
	// get the para
	command := c.GetString(ZLD_PARA_COMMAND) 
	fmt.Println("worker command=", command)
	zllogs.WriteDebugLog("GET request of worker page: command=%s", command)
	
	switch command {
	case ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case ZLD_CMD_LOAD_TASK:
		handleWorkerLoadTaskCmd(c)
	}		
}

func (c *WorkerController) Post() {
	// get the para
	command := c.GetString(ZLD_PARA_COMMAND) 
	fmt.Println("worker command=", command)
	zllogs.WriteDebugLog("Post request of worker page: command=%s", command)
	
	switch command {
	case ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case ZLD_CMD_LOAD_TASK:
		handleWorkerLoadTaskCmd(c)		
	}
}
