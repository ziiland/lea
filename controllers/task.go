package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lea/models"
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
	Page 				string
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
	task.TaskId = genTaskId("SHA001", "000A0001", "B")
	task.SponsorId = "ZLD00001"
	task.FarmId = "SHA001"
	task.CellId = "000A0001"
	task.PatchId = "B"
	task.CreateTime = time.Now().Unix()
	task.Type = 1;

	models.InsertTaskTableItem(task)
}

func (c *TaskController) Get() {
	// JUST FOR TEST
	//createWorkerTableItemForTest()

	// get the para
	workerId := c.GetString(ZLD_PARA_WORKERID)
	password := c.GetString(ZLD_PARA_PWD)

	item := new(LoginJsonData)
	item.Errcode = 1;
	// judgement the account
	if models.CheckWorkerLoginInfo(workerId, password) {
		// information correct!
		item.Errcode = 0;
	}

	item.Page = ""
	c.Data["json"] = item
	c.ServeJSON()	
}

func (c *TaskController) Post() {
	// get the para
	workerId := c.GetString(ZLD_PARA_WORKERID)
	password := c.GetString(ZLD_PARA_PWD)

	item := new(LoginJsonData)
	item.Errcode = 1;
	// judgement the account
	if models.CheckWorkerLoginInfo(workerId, password) {
		// information correct!
		item.Errcode = 0;
	}
		
	item.Page = ""
	c.Data["json"] = item
	c.ServeJSON()
}
