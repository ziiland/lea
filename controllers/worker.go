package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/models"
	"lea/zllogs"
	"strconv"
)


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type WorkerController struct {
	beego.Controller
}

// type WorkerJsonData struct {
// 	Tasks 				*[]models.ZldTaskData
// 	//Title               string
// 	Errcode             int
// }

type WorkerJsonData struct {
	Workers 			*[]models.ZldWorkerData
	Errcode				int
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
// func handleWorkerLoadTaskCmd(c *WorkerController) {
// 	workerId := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
// 	title := (c.GetSession(common.ZLD_PARA_TITLE)).(string)

// 	fmt.Println("handleWorkerLoadTaskCmd: worker=", workerId)
// 	if title == "员工" {
// 		fmt.Println("title=", title)
// 	} else if title == "经理" {
// 		fmt.Println("title=", title)
// 	}

// 	item := new(WorkerJsonData)
// 	//slice := make([]models.ZldTaskData, 1)
// 	if num, err := models.QueryMatchItemNums(workerId, "SHA001", title); err == nil {
// 		slice := make([]models.ZldTaskData, num)
// 		item.Tasks = &slice
// 		models.SelectTaskTableItemsWithFarmId(workerId, "SHA001", title, item.Tasks)
// 	}	
// 	item.Errcode = 0;

// 	c.Data["json"] = item
// 	c.ServeJSON()
// }

func loadOneWorkerInfo(c *WorkerController) {
	workerId := (c.GetSession(common.ZLD_PARA_WORKER)).(string)

	item := new(WorkerJsonData)
	slice := make([]models.ZldWorkerData, 1)
	item.Workers = &slice
	models.QueryWorkerTableItem(workerId, item.Workers)
	item.Errcode = 0

	c.Data["json"] = item
	c.ServeJSON()
}

func loadAllWorkersInfo(c *WorkerController) {
	workerId := (c.GetSession(common.ZLD_PARA_WORKER)).(string)

	item := new(WorkerJsonData)
	if num, err := models.QueryWorkerNumbers(workerId); err == nil {
		fmt.Println("loadAllWorkersInfo: num=", num);
		slice := make([]models.ZldWorkerData, num)
		item.Workers = &slice
		models.QueryAllWorkerTableItem(item.Workers)
	}	
	item.Errcode = 0;

	c.Data["json"] = item
	c.ServeJSON()	
}

func handleLoadWorkerInfo(c *WorkerController) {
	//workerId := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
	title := (c.GetSession(common.ZLD_PARA_TITLE)).(string)

	if title == "Admin" {
		// list all the workers
		loadAllWorkersInfo(c)
	} else {
		// only list self
		loadOneWorkerInfo(c)
	}
}
func stringToInt(s string) int64 {
	if value, err := strconv.ParseInt(s, 0, 0); err != nil {
		return 0
	} else {
		return value
	}
}
func handleAddWorkerInfo(c *WorkerController) {

	date := new(WorkerJsonData)
	item := models.NewZldWorkerDBData()

	item.WorkerId = c.GetString("WorkerId")
	item.Password =c.GetString("Password")
	item.Name =c.GetString("Name")
	item.Sex =c.GetString("Sex")
	item.IdentifyNo =c.GetString("IdentifyNo")
	item.Title =c.GetString("Title")
	item.CheckInTime = stringToInt(c.GetString("CheckInTime"))
	item.CheckOutTime = stringToInt(c.GetString("CheckInTime"))
	item.Comment =c.GetString("Comment")

	if !models.AlreadyHaveWorkerItem(item.WorkerId) {
		models.InsertWorkerTableItem(item)
		date.Errcode = 1;
		c.Data["json"] = date
		c.ServeJSON()

	}else{

		date.Errcode = 0;
		c.Data["json"] = date
		c.ServeJSON()
	}
}
///////////////////////////////////////////////////////////////////////////////
func (c *WorkerController) Get() {
	// get the para
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("worker command=", command)
	zllogs.WriteDebugLog("GET request of worker page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	// case common.ZLD_CMD_LOAD_TASK:
	// 	handleWorkerLoadTaskCmd(c)
	case common.ZLD_CMD_LOAD_WORKER:
		handleLoadWorkerInfo(c)
	case common.ZLD_CMD_ADD_WORKER:
		handleAddWorkerInfo(c)
	}		
}

func (c *WorkerController) Post() {
	// get the para
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("worker command=", command)
	zllogs.WriteDebugLog("Post request of worker page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	// case common.ZLD_CMD_LOAD_TASK:
	// 	handleWorkerLoadTaskCmd(c)
	case common.ZLD_CMD_LOAD_WORKER:
		handleLoadWorkerInfo(c)
	case common.ZLD_CMD_ADD_WORKER:
		handleAddWorkerInfo(c)
	}
}
