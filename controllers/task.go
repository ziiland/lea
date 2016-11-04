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
func queryMatchItemNums(workerId, title string) (int64, error){
	if title == common.ZLD_STR_WORKER {
		// only query task NOT CLOSED & worker/checker == workerId
		return models.QueryWorkerItemNums(workerId)
	} else {
		// query task NOT CLOSED
		return models.QueryAllOpenItemNums()
	}
}

func handleLoadTaskCmd(c *TaskController) {
	// load task condition: worker, time
	workerId := (c.GetSession(common.ZLD_PARA_WORKER)).(string)
	title := (c.GetSession(common.ZLD_PARA_TITLE)).(string)
	fmt.Printf("handleWorkerLoadTaskCmd: worker=%s, title=%s\n", workerId, title)

	// var v1 interface{} = etime
	// switch v1.(type) {
	// case int:
	// 	fmt.Println("etime is integer")
	// case string:
	// 	fmt.Println("etime is string")
	// default:
	// 	fmt.Println("etime is not integer or string")
	// }
	stime, _ := c.GetInt(common.ZLD_PARA_STIME)
	etime, _ := c.GetInt(common.ZLD_PARA_ETIME)	
	state := c.GetString(common.ZLD_PARA_STATE)
	worker := c.GetString(common.ZLD_PARA_WORKER)
	farm := c.GetString(common.ZLD_PARA_FARM)
	cell := c.GetString(common.ZLD_PARA_CELL)
	patch := c.GetString(common.ZLD_PARA_PATCH)

	fmt.Printf("stime=%d, etime=%d, state=%s, worker=%s farm=%s, cell=%s, patch=%s, title=%s\n", 
		stime, etime, state, worker, farm, cell, patch, title)	
	if title == common.ZLD_STR_WORKER {
		worker = workerId
	}
	item := new(TaskJsonData)
	//slice := make([]models.ZldTaskData, 1)
	// if num, err := queryMatchItemNums(worker, title); err == nil {
	// 	slice := make([]models.ZldTaskData, num)
	// 	item.Tasks = &slice
	// 	models.SelectTaskTableItemsWithFarmId(workerId, "SHA001", title, item.Tasks)
	// }	
	if tasks, err := models.SelectTaskTableItemWithConds(int64(stime), int64(etime),
	 worker, state, farm, cell, patch); err == nil {
	 	slice := make([]models.ZldTaskData, 0)
	 	for _, v := range tasks {
	 		slice = append(slice, v)
	 	}
	 	item.Tasks = &slice
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
