package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	"strconv"
	"time"
)

const ZLD_TASK_TBL_NAME string = "zld_task"

type ZldTaskData struct {
	Id				int32 		`orm:"pk;auto"`
	TaskId			string		// 
	SponsorId		string		// sponsor id: userid or workerid
	FarmId			string		// farm no
	PatchId			string		// 	
	WorkerId 		string
	CheckerId		string
	State 			string		// task state: doing, pending
	Type 			int64
	CreateTime 		int64
	StartTime 		int64
	EndTime 		int64
	CheckTime 		int64
	Score 			int64 
	UserComment		string
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldTaskDBData() *ZldTaskData{
	return &ZldTaskData{}
}

func CreateZldTaskTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_TASK_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `TaskId` varchar(32) NOT NULL DEFAULT '' COMMENT '任务号',", s)
	s = fmt.Sprintf("%s `SponsorId` varchar(10) NOT NULL DEFAULT '' COMMENT '发起人',", s)
	s = fmt.Sprintf("%s `FarmId` varchar(8) NOT NULL DEFAULT '' COMMENT '农场Id',", s)
	s = fmt.Sprintf("%s `PatchId` varchar(1) NOT NULL DEFAULT '' COMMENT '地块Id',", s)
	s = fmt.Sprintf("%s `WorkerId` varchar(10) NOT NULL DEFAULT '' COMMENT '施工员工',", s)
	s = fmt.Sprintf("%s `CheckerId` varchar(10) NOT NULL DEFAULT '' COMMENT '检查员工',", s)
	s = fmt.Sprintf("%s `State` varchar(32) NOT NULL DEFAULT '' COMMENT '任务状态',", s)
	s = fmt.Sprintf("%s `Type` int NOT NULL DEFAULT 0 COMMENT '任务种类',", s)	
	s = fmt.Sprintf("%s `CreateTime` int(10) NOT NULL DEFAULT 0 COMMENT '创建时间',", s)
	s = fmt.Sprintf("%s `StartTime` int(10) NOT NULL DEFAULT 0 COMMENT '开始时间',", s)
	s = fmt.Sprintf("%s `EndTime` int(10) NOT NULL DEFAULT 0 COMMENT '结束时间',", s)
	s = fmt.Sprintf("%s `CheckTime` int(10) NOT NULL DEFAULT 0 COMMENT '检查时间',", s)
	s = fmt.Sprintf("%s `Score` int(10) NOT NULL DEFAULT 0 COMMENT '评分',", s)
	s = fmt.Sprintf("%s `UserComment` text NOT NULL DEFAULT '' COMMENT '用户备注',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='任务表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_TASK_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_TASK_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_TASK_TBL_NAME)
	}	
}

func DoInsertTaskTableItem(item *ZldTaskData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_TASK_TBL_NAME)
	s = fmt.Sprintf("%s (`TaskId`, `SponsorId`, `FarmId`, `PatchId`, `WorkerId`, `CheckerId`, `State`, `Type`, `CreateTime`, `StartTime`, `EndTime`, `CheckTime`, `Score`, `UserComment`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%v', '%v', '%v', '%v', '%v', '%v', '%s', '%s');", s, item.TaskId, item.SponsorId, item.FarmId, item.PatchId, item.WorkerId, item.CheckerId, item.State, item.Type, time.Now().Unix(), item.StartTime, item.EndTime, item.CheckTime, item.Score, item.UserComment, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_TASK_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_TASK_TBL_NAME)
	}		
}

func AlreadyHaveTaskItem(id string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_TASK_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`TaskId` = '%s' );", s, id)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func SelectTaskTableItem(item *ZldTaskData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_TASK_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`TaskId` = '%s');", s, item.TaskId)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		ctime := (maps[0]["CreateTime"]).(string)
		item.CreateTime, _ = strconv.ParseInt(ctime, 10, 64)
		stime := (maps[0]["StartTime"]).(string)
		item.StartTime, _ = strconv.ParseInt(stime, 10, 64)
		etime := (maps[0]["EndTime"]).(string)
		item.EndTime, _ = strconv.ParseInt(etime, 10, 64)
		cktime := (maps[0]["CheckTime"]).(string)
		item.CheckTime, _ = strconv.ParseInt(cktime, 10, 64)

		ttype := (maps[0]["Type"]).(string)
		item.Type, _ =  strconv.ParseInt(ttype, 10, 64)
		score := (maps[0]["Score"]).(string)
		item.Score, _ = strconv.ParseInt(score, 10, 64)

		item.SponsorId = (maps[0]["SponsorId"]).(string)
		item.FarmId = (maps[0]["FarmId"]).(string)
		item.PatchId = (maps[0]["PatchId"]).(string)
		item.WorkerId = (maps[0]["WorkerId"]).(string)
		item.CheckerId = (maps[0]["CheckerId"]).(string)
		item.State = (maps[0]["State"]).(string)
		item.UserComment = (maps[0]["UserComment"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertTaskTableItem(item *ZldTaskData) {
	// first, try to create the table
	CreateZldTaskTable()

	// No same sn item, do insert
	if !AlreadyHaveTaskItem(item.TaskId) {
		DoInsertTaskTableItem(item)
	}
}

// func UpdateWorkerInfo(item *ZldWorkerData) error{
// 	s := fmt.Sprintf("UPDATE `%s`", ZLD_WORKER_TBL_NAME)
// 	s = fmt.Sprintf("%s SET ", s)
// 	s = fmt.Sprintf("%s `Password` = '%s',", s, item.Password)
// 	s = fmt.Sprintf("%s `Name` = '%s',", s, item.Name)
// 	s = fmt.Sprintf("%s `Sex` = '%s',", s, item.Sex)
// 	s = fmt.Sprintf("%s `IdentifyNo` = '%s'", s, item.IdentifyNo)
// 	s = fmt.Sprintf("%s `Title` = '%s'", s, item.Title)
// 	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, item.WorkerId)
// 	//fmt.Println("s=", s)

// 	o := orm.NewOrm()
// 	_, err := o.Raw(s).Exec()
// 	if err == nil {
// 		zllogs.WriteDebugLog("Update record(WorkerId=%s) in table %s  ...... DONE", item.WorkerId, ZLD_WORKER_TBL_NAME)
// 	} else {
// 		zllogs.WriteErrorLog("Update record(WorkerId=%s) in table %s  ...... ERROR", item.WorkerId, ZLD_WORKER_TBL_NAME)
// 	}

// 	return err
// }