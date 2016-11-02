package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	"strconv"
	//"time"
)

const ZLD_TASK_LOG_TBL_NAME string = "zld_task_log"

type ZldTaskLogData struct {
	Id				int32 		`orm:"pk;auto"`
	TaskId			string		// 
	Action			string
	OperatorId		string
	ActionTime		int64
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldTaskLogDBData() *ZldTaskLogData{
	return &ZldTaskLogData{}
}

func CreateZldTaskLogTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_TASK_LOG_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `TaskId` varchar(32) NOT NULL DEFAULT '' COMMENT '工作包号',", s)
	s = fmt.Sprintf("%s `Action` varchar(12) NOT NULL DEFAULT '' COMMENT '动作',", s)
	s = fmt.Sprintf("%s `OperatorId` varchar(10) NOT NULL DEFAULT '' COMMENT '工作员号',", s)
	s = fmt.Sprintf("%s `ActionTime` int(10) NOT NULL DEFAULT 0 COMMENT '动作时间',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='工作包日志表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_TASK_LOG_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_TASK_LOG_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_TASK_LOG_TBL_NAME)
	}	
}

func DoInsertTaskLogTableItem(item *ZldTaskLogData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_TASK_LOG_TBL_NAME)
	s = fmt.Sprintf("%s (`TaskId`, `Action`, `OperatorId`, `ActionTime`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', %v', '%s');", s, item.TaskId, item.Action, item.OperatorId, item.ActionTime, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_TASK_LOG_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_TASK_LOG_TBL_NAME)
	}		
}

func AlreadyHaveTaskLogItem(id string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_TASK_LOG_TBL_NAME)
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

// func SelectTaskTableItem(item *ZldTaskLogData) error{
// 	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_TASK_LOG_TBL_NAME)
// 	s = fmt.Sprintf("%s WHERE (`TaskId` = '%s');", s, item.TaskId)
// 	//fmt.Println("s=", s)

// 	var maps []orm.Params
// 	o := orm.NewOrm()
// 	num, err := o.Raw(s).Values(&maps)
// 	//fmt.Printf("num=%d, maps=%v\n", num, maps)

// 	if err == nil && num > 0 {
// 		ctime := (maps[0]["ActionTime"]).(string)
// 		item.CreateTime, _ = strconv.ParseInt(ctime, 10, 64)

// 		item.SponsorId = (maps[0]["OperatorId"]).(string)
// 		item.FarmId = (maps[0]["Action"]).(string)
// 		item.Comment = (maps[0]["Comment"]).(string)	
// 		fmt.Println("SelectItem=", *item)
// 	} else {
// 		fmt.Println("Select NONE! Error=%v", err)
// 	}

// 	return err	
// }

func InsertTaskLogTableItem(item *ZldTaskLogData) {
	// first, try to create the table
	CreateZldTaskLogTable()

	// No same sn item, do insert
	if !AlreadyHaveTaskLogItem(item.TaskId) {
		DoInsertTaskLogTableItem(item)
	}
}

func DecodeTaskLogOrmParamsToData(para orm.Params) (item ZldTaskLogData){
	item.TaskId = (para["TaskId"]).(string)
	id := (para["Id"]).(string)
	nId, _ := strconv.Atoi(id)
	item.Id = int32(nId)

	ctime := (para["ActionTime"]).(string)
	item.ActionTime, _ = strconv.ParseInt(ctime, 10, 64)

	item.Action = (para["Action"]).(string)
	item.OperatorId = (para["OperatorId"]).(string)
	item.Comment = (para["Comment"]).(string)
	return		
}

func QueryMatchLogItemNums(id string)(int64, error) {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_TASK_LOG_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`TaskId` = '%s');", s, id)	
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	return o.Raw(s).Values(&maps);
}

func SelectTaskLogTableItemsWithTaskId(taskid string, logs *ZldTaskLogData)(num int64, err error) {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_TASK_LOG_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`TaskId` = '%s');", s, taskid)
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err = o.Raw(s).Values(&maps);

	if err == nil && num > 0 {
		*logs = DecodeTaskLogOrmParamsToData(maps[0])
	}

	fmt.Println("task log=", logs)
	return num, err
}