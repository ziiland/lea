package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	"strconv"
	"time"
)

const ZLD_WORKER_TBL_NAME string = "zld_worker"

type ZldWorkerData struct {
	Id				int32 		`orm:"pk;auto"`
	WorkerId		string		// working id
	Password		string		// password
	Name			string		// 
	Sex				string		// 
	IdentifyNo		string		//
	Title			string		// Admin/PM/PE
	//Farm            string      // belong to which farm
	CheckInTime		int64
	CheckOutTime	int64
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldWorkerDBData() *ZldWorkerData{
	return &ZldWorkerData{}
}

func CreateZldWorkerTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `WorkerId` varchar(10) NOT NULL DEFAULT '' COMMENT '工号',", s)
	s = fmt.Sprintf("%s `Password` varchar(16) NOT NULL DEFAULT '' COMMENT '密码',", s)
	s = fmt.Sprintf("%s `Name` varchar(128) NOT NULL DEFAULT '' COMMENT '中文名',", s)
	s = fmt.Sprintf("%s `Sex` varchar(8) NOT NULL DEFAULT '' COMMENT '性别',", s)
	s = fmt.Sprintf("%s `IdentifyNo` varchar(32) NOT NULL DEFAULT '' COMMENT '身份证号',", s)
	s = fmt.Sprintf("%s `Title` varchar(32) NOT NULL DEFAULT '' COMMENT '角色',", s)
	s = fmt.Sprintf("%s `CheckInTime` int(10) NOT NULL DEFAULT 0 COMMENT '入职时间',", s)
	s = fmt.Sprintf("%s `CheckOutTime` int(10) NOT NULL DEFAULT 0 COMMENT '离职时间',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='员工表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_WORKER_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_WORKER_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_WORKER_TBL_NAME)
	}	
}

func DoInsertWorkerTableItem(item *ZldWorkerData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s (`WorkerId`, `Password`, `Name`, `Sex`, `IdentifyNo`, `Title`, `CheckInTime`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s', '%v', '%s');", s, item.WorkerId, item.Password, item.Name, item.Sex, item.IdentifyNo, item.Title, time.Now().Unix(), item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_WORKER_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_WORKER_TBL_NAME)
	}		
}

func AlreadyHaveWorkerItem(id string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s' );", s, id)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func CheckWorkerLoginInfo(workerId, pwd string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s' AND `Password` = '%s' AND `CheckOutTime` = 0);", s, workerId, pwd)
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num == 1 {
		return true
	} else {
		return false
	}
}

func SelectWorkerTableItem(item *ZldWorkerData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, item.WorkerId)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		stime := (maps[0]["CheckInTime"]).(string)
		item.CheckInTime, _ = strconv.ParseInt(stime, 10, 64)
		etime := (maps[0]["CheckOutTime"]).(string)
		item.CheckOutTime, _ = strconv.ParseInt(etime, 10, 64)
		item.Password = (maps[0]["Password"]).(string)
		item.Name = (maps[0]["Name"]).(string)
		item.Sex = (maps[0]["Sex"]).(string)
		item.IdentifyNo = (maps[0]["IdentifyNo"]).(string)
		item.Title = (maps[0]["Title"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertWorkerTableItem(item *ZldWorkerData) {
	// first, try to create the table
	CreateZldWorkerTable()

	// No same sn item, do insert
	if !AlreadyHaveWorkerItem(item.WorkerId) {
		DoInsertWorkerTableItem(item)
	}
}

func UpdateWorkerPassword(id, password string) error {
	s := fmt.Sprintf("UPDATE `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s SET ", s)
	s = fmt.Sprintf("%s `Password` = '%v'", s, password)
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, id)	
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		zllogs.WriteDebugLog("Update password (WorkerId=%s) ...... SUCCESS!", id)
	} else {
		zllogs.WriteErrorLog("Update password (WorkerId=%s) ...... ERROR!", id)
	}

	return err	
}

func UpdateWorkerTitle(id, title string) error{
	s := fmt.Sprintf("UPDATE `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s SET ", s)
	s = fmt.Sprintf("%s `Title` = '%v'", s, title)
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, id)	
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		zllogs.WriteDebugLog("Update Title (WorkerId=%s) ...... SUCCESS!", id)
	} else {
		zllogs.WriteErrorLog("Update Title (WorkerId=%s) ...... ERROR!", id)
	}

	return err		
}

func UpdateWorkerInfo(item *ZldWorkerData) error{
	s := fmt.Sprintf("UPDATE `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s SET ", s)
	s = fmt.Sprintf("%s `Password` = '%s',", s, item.Password)
	s = fmt.Sprintf("%s `Name` = '%s',", s, item.Name)
	s = fmt.Sprintf("%s `Sex` = '%s',", s, item.Sex)
	s = fmt.Sprintf("%s `IdentifyNo` = '%s',", s, item.IdentifyNo)
	s = fmt.Sprintf("%s `Title` = '%s',", s, item.Title)
	s = fmt.Sprintf("%s `Comment` = '%s'", s, item.Comment)
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, item.WorkerId)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		zllogs.WriteDebugLog("Update record(WorkerId=%s) in table %s  ...... DONE", item.WorkerId, ZLD_WORKER_TBL_NAME)
	} else {
		zllogs.WriteErrorLog("Update record(WorkerId=%s) in table %s  ...... ERROR", item.WorkerId, ZLD_WORKER_TBL_NAME)
	}

	return err
}

func UpdateWorkerCheckOutTime(id string) error {
	s := fmt.Sprintf("UPDATE `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s SET ", s)
	s = fmt.Sprintf("%s `CheckOutTime` = '%v'", s, time.Now().Unix())
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, id)	

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		zllogs.WriteDebugLog("Update record(WorkerId=%s) in table %s  ...... DONE", id, ZLD_WORKER_TBL_NAME)
	} else {
		zllogs.WriteErrorLog("Update record(WorkerId=%s) in table %s  ...... ERROR", id, ZLD_WORKER_TBL_NAME)
	}

	return err	
}

func DecodeWorkerOrmParamsToData(para orm.Params) (item ZldWorkerData){
	id := (para["Id"]).(string)
	nId, _ := strconv.Atoi(id)
	item.Id = int32(nId)

	ctime := (para["CheckInTime"]).(string)
	item.CheckInTime, _ = strconv.ParseInt(ctime, 10, 64)
	stime := (para["CheckOutTime"]).(string)
	item.CheckOutTime, _ = strconv.ParseInt(stime, 10, 64)

	item.WorkerId = (para["WorkerId"]).(string)
	item.Password = (para["Password"]).(string)
	item.Name = (para["Name"]).(string)
	item.Sex = (para["Sex"]).(string)
	item.IdentifyNo = (para["IdentifyNo"]).(string)
	item.Title = (para["Title"]).(string)
	item.Comment = (para["Comment"]).(string)

	return		
}

func QueryInSvcWorkerNumbers(id string) (int64, error) {
	s := fmt.Sprintf("SELECT * FROM `%s` WHERE (`CheckOutTime` = 0);", ZLD_WORKER_TBL_NAME)
	//s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, id)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	return  o.Raw(s).Values(&maps)
}

func QueryWorkerTableItem(id string) (ZldWorkerData, error){
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, id)
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	var worker ZldWorkerData
	if err == nil && num > 0 {
		worker = DecodeWorkerOrmParamsToData(maps[0])
	}
	fmt.Println("worker=", worker)

	return worker, err
}

func QueryInSvcAllWorkerTableItem() ([]ZldWorkerData, error) {
	s := fmt.Sprintf("SELECT * FROM `%s` WHERE (`CheckOutTime` = 0);", ZLD_WORKER_TBL_NAME)
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	workers := make([]ZldWorkerData, num)
	if err == nil && num > 0 {
		for i, v := range maps {
			workers[i] = DecodeWorkerOrmParamsToData(v)
		}
	}
	fmt.Println("workers=", workers)
	return workers, err
}

func QueryAllWorkersTableItem() ([]ZldWorkerData, error) {
	s := fmt.Sprintf("SELECT * FROM `%s`;", ZLD_WORKER_TBL_NAME)
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	workers := make([]ZldWorkerData, num)
	if err == nil && num > 0 {
		for i, v := range maps {
			workers[i] = DecodeWorkerOrmParamsToData(v)
		}
	}
	fmt.Println("workers=", workers)
	return workers, err
}