package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/models"
	"lea/zllogs"
	"strconv"
	"time"
)


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type TaskController struct {
	beego.Controller
}

// type TaskJsonData struct {
// 	Page 				string
// 	//Title               string
// 	Errcode             int
// }

type TaskJsonData struct {
	Tasks 				*[]models.ZldTaskData
	//Title               string
	Errcode             int
}

///////////////////////////////////////////////////////////////////////////////
func genTaskId(farm, cell, patch string) string{
	id := farm
	id += cell
	id += patch

	t := time.Now().Unix()
	id += strconv.FormatInt(t, 10)

	var index int = 0
	for{
		s := fmt.Sprintf("%07d", index)
		if models.AlreadyHaveTaskItem(id + s) {
			index += 10
		} else {
			id += s
			break;
		}
	}
	
	fmt.Printf("id=%s\n", id)
	return id
}

func createTaskTableItemForTest() {
	// create some tasks
	task := models.NewZldTaskDBData()
	task.SponsorId = "ZLD00001"
	task.FarmId = "SHA001"
	task.CellId = "000A0001"
	task.PatchId = "B"
	task.CreateTime = time.Now().Unix()
	task.Type = 1;
	task.TaskId = genTaskId("SHA001", "000A0001", "B")
	models.InsertTaskTableItem(task)

	task.TaskId = genTaskId("SHA001", "000A0001", "B")
	models.InsertTaskTableItem(task)

	task.TaskId = genTaskId("SHA001", "000A0001", "B")
	models.InsertTaskTableItem(task)
}

///////////////////////////////////////////////////////////////////////////////
func handleLoadTaskCmd(c *TaskController) {
	// load task condition: worker, time
	workerId := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
	title := (c.GetSession(common.ZLD_PARA_TITLE)).(string)

	// base the worker, get the farmers/title

	fmt.Println("handleWorkerLoadTaskCmd: worker=", workerId)
	if title == "员工" {
		fmt.Println("title=", title)
	} else if title == "经理" {
		fmt.Println("title=", title)
	}

	item := new(TaskJsonData)
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

func handleQuerytaskCmd(c *TaskController) {
	data := new(TaskJsonData)
	slice := make([]models.ZldTaskData, 1)
	data.Tasks = &slice

	taskId := c.GetString(common.ZLD_PARA_TASKID) 
	data.Errcode = 1;
	if _, err := models.SelectTaskTableItemsWithTaskId(taskId, &slice[0]); err == nil {
		data.Errcode = 0;
	}	

	c.Data["json"] = data
	c.ServeJSON()
}

///////////////////////////////////////////////////////////////////////////////
func (c *TaskController) Get() {
	// get the para
	// get the para
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("task command=", command)
	zllogs.WriteDebugLog("GET request of task page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_LOAD_TASK:
		handleLoadTaskCmd(c)
	case common.ZLD_CMD_QUERY_TASK:
		handleQuerytaskCmd(c)
	}	
}

func (c *TaskController) Post() {
	// get the para
	command := c.GetString(common.ZLD_PARA_COMMAND) 
	fmt.Println("task command=", command)
	zllogs.WriteDebugLog("POST request of task page: command=%s", command)
	
	switch command {
	case common.ZLD_CMD_LOAD_PARA:
		handleLoadParaCmd(&c.Controller)
	case common.ZLD_CMD_UNLOAD:
		handleUnloadCmd(&c.Controller)
	case common.ZLD_CMD_LOAD_TASK:
		handleLoadTaskCmd(c)
	case common.ZLD_CMD_QUERY_TASK:
		handleQuerytaskCmd(c)
	}	
}
