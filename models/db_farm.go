package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	"strconv"
	//"time"
)

const ZLD_FARM_TBL_NAME string = "zld_farm"

type ZldFarmData struct {
	Id				int32 		`orm:"pk;auto"`
	FarmId			string		// farm id
	City			string
	District		string
	Village			string
	Longitude		float64
	Latitude		float64
	RentTime		int64
	RentMonths		int32
	Tenant			string
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldFarmDBData() *ZldFarmData{
	return &ZldFarmData{}
}

func CreateZldFarmTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_FARM_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `FarmId` varchar(10) NOT NULL DEFAULT '' COMMENT '农场号',", s)
	s = fmt.Sprintf("%s `City` varchar(16) NOT NULL DEFAULT '' COMMENT '城市',", s)
	s = fmt.Sprintf("%s `District` varchar(128) NOT NULL DEFAULT '' COMMENT '区县',", s)
	s = fmt.Sprintf("%s `Village` varchar(8) NOT NULL DEFAULT '' COMMENT '村',", s)
	s = fmt.Sprintf("%s `Longitude` float NOT NULL DEFAULT 0.0 COMMENT '经度',", s)
	s = fmt.Sprintf("%s `Latitude` float NOT NULL DEFAULT 0.0 COMMENT '维度',", s)
	s = fmt.Sprintf("%s `RentTime` int(10) NOT NULL DEFAULT 0 COMMENT '起租时间',", s)
	s = fmt.Sprintf("%s `RentMonths` int NOT NULL DEFAULT 0 COMMENT '租期',", s)
	s = fmt.Sprintf("%s `Tenant` varchar(256) NOT NULL DEFAULT '' COMMENT '承租人',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='农场表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_FARM_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_FARM_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_FARM_TBL_NAME)
	}	
}

func DoInsertFarmTableItem(item *ZldFarmData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_FARM_TBL_NAME)
	s = fmt.Sprintf("%s (`FarmId`, `City`, `District`, `Village`, `Longitude`, `Latitude`, `RentTime`, `RentMonths`, `Tenant`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%v', '%v', '%v', '%v', '%s', '%s');", s, item.FarmId, item.City, item.District, item.Village, item.Longitude, item.Latitude, item.RentTime, item.RentMonths, item.Tenant, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_FARM_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_FARM_TBL_NAME)
	}		
}

func AlreadyHaveFarmItem(id string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_FARM_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s' );", s, id)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func SelectFarmTableItem(item *ZldFarmData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_FARM_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s');", s, item.FarmId)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		stime := (maps[0]["RentTime"]).(string)
		item.RentTime, _ = strconv.ParseInt(stime, 10, 64)
		month := (maps[0]["RentMonths"]).(string)
		rentMonth, _ := strconv.Atoi(month)
		item.RentMonths = int32(rentMonth)
		longitude := (maps[0]["Longitude"]).(string)
		item.Longitude, _ = strconv.ParseFloat(longitude, 64)
		latitude := (maps[0]["Latitude"]).(string)
		item.Latitude, _ = strconv.ParseFloat(latitude, 64)

		item.City = (maps[0]["City"]).(string)
		item.District = (maps[0]["District"]).(string)
		item.Village = (maps[0]["Village"]).(string)
		item.Tenant = (maps[0]["Tenant"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertFarmTableItem(item *ZldFarmData) {
	// first, try to create the table
	CreateZldFarmTable()

	// No same sn item, do insert
	if !AlreadyHaveFarmItem(item.FarmId) {
		DoInsertFarmTableItem(item)
	}
}

// func UpdateWorkerPassword(id, password string) error {
// 	s := fmt.Sprintf("UPDATE `%s`", ZLD_WORKER_TBL_NAME)
// 	s = fmt.Sprintf("%s SET ", s)
// 	s = fmt.Sprintf("%s `Password` = '%v',", s, password)
// 	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, id)	
// 	//fmt.Println("s=", s)

// 	o := orm.NewOrm()
// 	_, err := o.Raw(s).Exec()
// 	if err == nil {
// 		zllogs.WriteDebugLog("Update password (WorkerId=%s) ...... SUCCESS!", id)
// 	} else {
// 		zllogs.WriteErrorLog("Update password (WorkerId=%s) ...... ERROR!", id)
// 	}

// 	return err	
// }

// func UpdateWorkerTitle(id, title string) error{
// 	s := fmt.Sprintf("UPDATE `%s`", ZLD_WORKER_TBL_NAME)
// 	s = fmt.Sprintf("%s SET ", s)
// 	s = fmt.Sprintf("%s `Title` = '%v',", s, title)
// 	s = fmt.Sprintf("%s WHERE (`WorkerId` = '%s');", s, id)	
// 	//fmt.Println("s=", s)

// 	o := orm.NewOrm()
// 	_, err := o.Raw(s).Exec()
// 	if err == nil {
// 		zllogs.WriteDebugLog("Update Title (WorkerId=%s) ...... SUCCESS!", id)
// 	} else {
// 		zllogs.WriteErrorLog("Update Title (WorkerId=%s) ...... ERROR!", id)
// 	}

// 	return err		
// }

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