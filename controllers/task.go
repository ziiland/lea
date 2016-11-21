package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/common"
	"lea/models"
	"lea/utils/simplejson"
	"lea/zllogs"
	"strings"
	"strconv"
	"time"
)


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type TaskController struct {
	beego.Controller
}

type TaskJsonData struct {
	Tasks 				*[]models.ZldTaskData
	//Title               string
	Errcode             int
}

type TaskLogJsonData struct {
	Logs 				*[]models.ZldTaskLogData
	Errcode 			int
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
	task.TaskId = genTaskId("SHA01A", "000A0001", "B")
	models.InsertTaskTableItem(task)

	task.TaskId = genTaskId("SHA01A", "000A0001", "B")
	models.InsertTaskTableItem(task)

	task.TaskId = genTaskId("SHA01A", "000A0001", "B")
	models.InsertTaskTableItem(task)
}

func createTaskLogTableItemForTest() {
	log := models.NewZldTaskLogDBData()
	log.TaskId = "SHA01A000A0001B14768647080000000"
	log.Action = "Create" 
	log.OperatorId = "williamzhang"
	log.ActionTime = time.Now().Unix() - 367890
	models.InsertTaskLogTableItem(log)

	log.Action = "Assign"
	log.OperatorId = "ZLD00001"
	log.ActionTime = time.Now().Unix() - 95678
	models.InsertTaskLogTableItem(log)
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

	fmt.Printf("stime=%d, etime=%d, state=%s, worker=%s, farm=%s, cell=%s, patch=%s, title=%s\n", 
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
	item.Errcode = 1	
	if tasks, err := models.SelectTaskTableItemWithConds(int64(stime), int64(etime),
	 worker, state, farm, cell, patch); err == nil {
	 	slice := make([]models.ZldTaskData, 0)
	 	for _, v := range tasks {
	 		slice = append(slice, v)
	 	}
	 	item.Tasks = &slice
	 	item.Errcode = 0
	}	

	c.Data["json"] = item
	c.ServeJSON()
}

func handleAssignTaskCmd(c *TaskController) {
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)

	item := new(TaskJsonData)
	item.Errcode = 1
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	if tasks, err := paraJSON.Get("Tasks").StringArray(); err == nil {
		worker, _ := paraJSON.Get("Worker").String()
		checker, _ := paraJSON.Get("Checker").String()

		models.AssignTasksItem(tasks, worker, checker)
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()
}

func handleCheckTaskCmd(c *TaskController) {
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)

	item := new(TaskJsonData)
	item.Errcode = 1	

	// para
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	fmt.Println("paraJSON=", paraJSON)
	if tasks, err := paraJSON.Get("Tasks").StringArray(); err == nil {
		fmt.Println("tasks=", tasks)
		models.CheckTasksItem(tasks)
		item.Errcode = 0
	}
	c.Data["json"] = item
	c.ServeJSON()	
}

func handleCancelTaskCmd(c *TaskController) {
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)

	item := new(TaskJsonData)
	item.Errcode = 1
	// para
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	fmt.Println("paraJSON=", paraJSON)
	if tasks, err := paraJSON.Get("Tasks").StringArray(); err == nil {
		fmt.Println("tasks=", tasks)
		models.CancelTasksItem(tasks)
		item.Errcode = 0
	}
	c.Data["json"] = item
	c.ServeJSON()
}

func handleQueryTaskCmd(c *TaskController) {
	taskId := c.GetString(common.ZLD_PARA_TASKID)

	item := new(TaskLogJsonData) 
	item.Errcode = 1;
	if logs, err := models.SelectTaskZLogTableItemsWithTaskId(taskId); err == nil {
		slice := make([]models.ZldTaskLogData, 0)
		for _, v := range(logs) {
			slice = append(slice, v)
		}
		item.Logs = &slice
		item.Errcode = 0;
	}	

	c.Data["json"] = item
	c.ServeJSON()
}

func handleCloseTaskCmd(c *TaskController) {
	strCmdPara := c.GetString("CmdPara")
	bytesCmdPara := []byte(strCmdPara)
	fmt.Println("bytesCmdPara=", bytesCmdPara)

	item := new(TaskJsonData)
	item.Errcode = 1
	// para
	paraJSON, _ := simplejson.NewJson(bytesCmdPara)
	fmt.Println("paraJSON=", paraJSON)
	if tasks, err := paraJSON.Get("Tasks").StringArray(); err == nil {
		fmt.Println("tasks=", tasks)
		models.CloseTasksItem(tasks)
		item.Errcode = 0
	}
	c.Data["json"] = item
	c.ServeJSON()
}

func handleArchiveTaskCmd(c *TaskController) {
	taskId := c.GetString(common.ZLD_PARA_TASKID)

	item := new(TaskJsonData)
	item.Errcode = 1

	// get data
	if task, err := models.SelectTaskTableItemsWithTaskId(taskId); err == nil {
		// add to task archive table
		models.InsertTaskArchivedTableItem(task)
		models.DeleteTaskItem(taskId)
		item.Errcode = 0
	}	

	c.Data["json"] = item
	c.ServeJSON()
}

func handleAddTaskCmd(c *TaskController) {
	item := new(TaskJsonData)

	task := models.NewZldTaskDBData()
	task.SponsorId = strings.ToUpper((c.GetSession(common.ZLD_PARA_WORKER)).(string))
	task.FarmId = c.GetString(common.ZLD_PARA_FARM)
	task.CellId = c.GetString(common.ZLD_PARA_CELL)
	task.PatchId = c.GetString(common.ZLD_PARA_PATCH)
	task.CreateTime = time.Now().Unix()
	task.WorkerId =  c.GetString(common.ZLD_PARA_WORKER)
	ctype := c.GetString(common.ZLD_PARA_TYPE)
	task.Type, _ = strconv.ParseInt(ctype, 10, 64)
	task.TaskId = genTaskId(task.FarmId, task.CellId, task.PatchId)
	task.Comment = c.GetString(common.ZLD_PARA_COMMENT)
	fmt.Println("AddTask: task=", task)
	models.InsertTaskTableItem(task)

	item.Errcode = 0
	c.Data["json"] = item
	c.ServeJSON()
}

func handleBeginTaskCmd(c *TaskController) {
	taskId := c.GetString(common.ZLD_PARA_TASKID)

	item := new(TaskJsonData)
	item.Errcode = 1

	if _, err := models.BeginTaskItem(taskId); err == nil {
		item.Errcode = 0
	}
	c.Data["json"] = item
	c.ServeJSON()	
}

func handleSubmitTaskCmd(c *TaskController) {
	taskId := c.GetString(common.ZLD_PARA_TASKID)
	item := new(TaskJsonData)
	item.Errcode = 1

	if _, err := models.SubmitTaskItem(taskId); err == nil {
		item.Errcode = 0
	}

	c.Data["json"] = item
	c.ServeJSON()	
}

///////////////////////////////////////////////////////////////////////////////
func (c *TaskController) Get() {
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
	case common.ZLD_CMD_ASSIGN_TASK:
		handleAssignTaskCmd(c)
	case common.ZLD_CMD_CHECK_TASK:
		handleCheckTaskCmd(c)
	case common.ZLD_CMD_QUERY_TASK:
		handleQueryTaskCmd(c)
	case common.ZLD_CMD_ARCHIVE_TASK:
		handleArchiveTaskCmd(c)
	case common.ZLD_CMD_CANCEL_TASK:
		handleCancelTaskCmd(c)
	case common.ZLD_CMD_CLOSE_TASK:
		handleCloseTaskCmd(c)
	case common.ZLD_CMD_ADD_TASK:
		handleAddTaskCmd(c)
	case common.ZLD_CMD_SUBMIT_TASK:
		handleSubmitTaskCmd(c)
	case common.ZLD_CMD_BEGIN_TASK:
		handleBeginTaskCmd(c)
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
	case common.ZLD_CMD_ASSIGN_TASK:
		handleAssignTaskCmd(c)
	case common.ZLD_CMD_CHECK_TASK:
		handleCheckTaskCmd(c)		
	case common.ZLD_CMD_QUERY_TASK:
		handleQueryTaskCmd(c)
	case common.ZLD_CMD_ARCHIVE_TASK:
		handleArchiveTaskCmd(c)	
	case common.ZLD_CMD_CANCEL_TASK:
		handleCancelTaskCmd(c)
	case common.ZLD_CMD_CLOSE_TASK:
		handleCloseTaskCmd(c)
	case common.ZLD_CMD_ADD_TASK:
		handleAddTaskCmd(c)
	case common.ZLD_CMD_SUBMIT_TASK:
		handleSubmitTaskCmd(c)
	case common.ZLD_CMD_BEGIN_TASK:
		handleBeginTaskCmd(c)
	}	
}
