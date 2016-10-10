package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
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
	s = fmt.Sprintf("%s `WorkerId` varchar(32) NOT NULL DEFAULT '' COMMENT 'GH',", s)
	s = fmt.Sprintf("%s `Password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',", s)
	s = fmt.Sprintf("%s `Name` varchar(128) NOT NULL DEFAULT '' COMMENT '中文名',", s)
	s = fmt.Sprintf("%s `Sex` varchar(32) NOT NULL DEFAULT '' COMMENT 'xb',", s)
	s = fmt.Sprintf("%s `IdentifyNo` varchar(32) NOT NULL DEFAULT '' COMMENT 'sfz',", s)
	s = fmt.Sprintf("%s `Title` varchar(32) NOT NULL DEFAULT '' COMMENT '角色',", s)
	s = fmt.Sprintf("%s `CheckInTime` int(10) NOT NULL DEFAULT 0 COMMENT 'rzsj',", s)
	s = fmt.Sprintf("%s `CheckOutTime` int(10) NOT NULL DEFAULT 0 COMMENT 'lzsj',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';", s)
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

func DoInsertZldWorkerTableItem(item *ZldWorkerData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_WORKER_TBL_NAME)
	s = fmt.Sprintf("%s (`WorkerId`, `Password`, `Name`, `Sex`, `IdentifyNo`, `Title`, `CheckInTime`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s', '%v');", s, item.WorkerId, item.Password, item.Name, item.Sex, item.IdentifyNo, item.Title, time.Now().Unix())
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
