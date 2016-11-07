package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/models"
	"lea/zllogs"
	//"strconv"
)


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type WorkerController struct {
	beego.Controller
}

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
	item.Errcode = 1

	if worker, err := models.QueryWorkerTableItem(workerId); err == nil {
		slice := make([]models.ZldWorkerData, 1)
		slice[0] = worker
		item.Workers = &slice		
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()
}

func loadAllInSvcWorkersInfo(c *WorkerController) {
	item := new(WorkerJsonData)
	item.Errcode = 1;
	if workers, err := models.QueryInSvcAllWorkerTableItem(); err == nil {
		slice := make([]models.ZldWorkerData, 0)
		for _, v := range workers {
			slice = append(slice, v)
		}
		item.Workers = &slice
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()		
}

func loadAllWorkersInfo(c *WorkerController) {
	item := new(WorkerJsonData)
	item.Errcode = 1
	if workers, err := models.QueryAllWorkersTableItem(); err == nil {
		slice := make([]models.ZldWorkerData, 0)
		for _, v := range workers {
			slice = append(slice, v)
		}
		item.Workers = &slice
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()	
}

func handleLoadWorkerInfo(c *WorkerController) {
	//workerId := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
	title := (c.GetSession(common.ZLD_PARA_TITLE)).(string)

	if title == common.ZLD_STR_ADMIN {
		// list all the workers
		loadAllWorkersInfo(c)
	} else if title == common.ZLD_STR_MANAGER{
		// list all in service workers
		loadAllInSvcWorkersInfo(c)
	} else {
		// only list self
		loadOneWorkerInfo(c)
	}
}

func handleChangePassword(c *WorkerController) {
	worker := c.GetString(common.ZLD_PARA_WORKER)
	password := c.GetString(common.ZLD_PARA_PWD)

	data := new(WorkerJsonData)
	data.Errcode = 1
	if err := models.UpdateWorkerPassword(worker, password); err == nil {
		data.Errcode = 0
	}
	c.Data["json"] = data
	c.ServeJSON()	
}

func handleUpdateWorkerInfo(c *WorkerController) {
	item := models.NewZldWorkerDBData()
	item.WorkerId = c.GetString(common.ZLD_PARA_WORKER)
	item.Password = c.GetString(common.ZLD_PARA_PWD)
	item.Name = c.GetString(common.ZLD_PARA_NAME)
	item.Sex = c.GetString(common.ZLD_PARA_SEX)
	item.IdentifyNo = c.GetString(common.ZLD_PARA_ID)
	item.Title = c.GetString(common.ZLD_PARA_TITLE)
	item.Comment = c.GetString(common.ZLD_PARA_COMMENT)

	data := new(WorkerJsonData)
	data.Errcode = 1
	if err := models.UpdateWorkerInfo(item); err == nil {
		data.Errcode = 0
	}
	c.Data["json"] = data
	c.ServeJSON()			
}

func handleAddWorkerInfo(c *WorkerController) {
	data := new(WorkerJsonData)
	item := models.NewZldWorkerDBData()

	item.WorkerId = c.GetString(common.ZLD_PARA_WORKER)
	item.Password = c.GetString(common.ZLD_PARA_PWD)
	item.Name = c.GetString(common.ZLD_PARA_NAME)
	item.Sex = c.GetString(common.ZLD_PARA_SEX)
	item.IdentifyNo = c.GetString(common.ZLD_PARA_ID)
	item.Title = c.GetString(common.ZLD_PARA_TITLE)
	item.Comment = c.GetString(common.ZLD_PARA_COMMENT)

	if !models.AlreadyHaveWorkerItem(item.WorkerId) {
		models.InsertWorkerTableItem(item)
		data.Errcode = 0
	}else{
		data.Errcode = 1
	}
	c.Data["json"] = data
	c.ServeJSON()	
}

func handleDelWorkerCmd(c *WorkerController) {
	data := new(WorkerJsonData)
	workerId := c.GetString(common.ZLD_PARA_WORKER)

	data.Errcode = 1;
	if err := models.UpdateWorkerCheckOutTime(workerId); err == nil{
		data.Errcode = 0
	}

	c.Data["json"] = data
	c.ServeJSON()	
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
	case common.ZLD_CMD_LOAD_WORKER:
		handleLoadWorkerInfo(c)
	case common.ZLD_CMD_ADD_WORKER:
		handleAddWorkerInfo(c)
	case common.ZLD_CMD_DEL_WORKER:
	 	handleDelWorkerCmd(c)
	case common.ZLD_CMD_CHGPWD_WORKER:
	 	handleChangePassword(c)
	case common.ZLD_CMD_UPD_WORKER:
		handleUpdateWorkerInfo(c)
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
	case common.ZLD_CMD_LOAD_WORKER:
		handleLoadWorkerInfo(c)
	case common.ZLD_CMD_ADD_WORKER:
		handleAddWorkerInfo(c)
	case common.ZLD_CMD_DEL_WORKER:
		handleDelWorkerCmd(c)
	case common.ZLD_CMD_CHGPWD_WORKER:
	 	handleChangePassword(c)	
	case common.ZLD_CMD_UPD_WORKER:
		handleUpdateWorkerInfo(c)	 		
	}
}
