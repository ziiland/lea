package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/utils/simplejson"
	"lea/zllogs"
	"strconv"
	//"time"
)

const ZLD_PACKET_TBL_NAME string = "zld_packet"

type ZldPacketData struct {
	Id				int32 		`orm:"pk;auto"`
	TaskId          string
	Sender			string		
	Receiver		string      
	Address			string
	ExpressNo		string
	RMobile			string
	SendTime		int64
	ExpressFee		float64
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldPacketDBData() *ZldPacketData{
	return &ZldPacketData{}
}

func CreateZldPacketTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_PACKET_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `TaskId` varchar(32) NOT NULL DEFAULT '' COMMENT '工作包号',", s)
	s = fmt.Sprintf("%s `Sender` varchar(128) NOT NULL DEFAULT '' COMMENT '发件人',", s)
	s = fmt.Sprintf("%s `Receiver` varchar(128) NOT NULL DEFAULT '' COMMENT '收件人',", s)
	s = fmt.Sprintf("%s `Address` varchar(256) NOT NULL DEFAULT '' COMMENT '地址',", s)
	s = fmt.Sprintf("%s `ExpressNo` varchar(32) NOT NULL DEFAULT '' COMMENT '快递单号',", s)
	s = fmt.Sprintf("%s `RMobile` varchar(16) NOT NULL DEFAULT '' COMMENT '收件人手机号',", s)	
	s = fmt.Sprintf("%s `SendTime` int(10) NOT NULL DEFAULT 0 COMMENT '寄出时间',", s)
	s = fmt.Sprintf("%s `ExpressFee` float NOT NULL DEFAULT 0.0 COMMENT '快递费',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='包裹表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_PACKET_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_PACKET_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_PACKET_TBL_NAME)
	}	
}

func DoInsertPacketTableItem(item *ZldPacketData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_PACKET_TBL_NAME)
	s = fmt.Sprintf("%s (`TaskId`, `Sender`, `Receiver`, `Address`, `ExpressNo`, `RMobile`, `SendTime`, `ExpressFee`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s', '%v', '%v', '%s');", s, item.TaskId, item.Sender, item.Receiver, item.Address, item.ExpressNo, item.RMobile, item.SendTime, item.ExpressFee, item.Comment)
	fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_PACKET_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_PACKET_TBL_NAME)
	}		
}

func AlreadyHavePacketItem(ExpressNo string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PACKET_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`ExpressNo` = '%s');", s, ExpressNo)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func SelectPacketTableItem(item *ZldPacketData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PACKET_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`ExpressNo` = '%s');", s, item.ExpressNo)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		stime := (maps[0]["SendTime"]).(string)
		item.SendTime, _ = strconv.ParseInt(stime, 10, 64)
		fee := (maps[0]["ExpressFee"]).(string)
		item.ExpressFee, _ = strconv.ParseFloat(fee, 64)

		item.Sender = (maps[0]["Sender"]).(string)
		item.Receiver = (maps[0]["Receiver"]).(string)
		item.Address = (maps[0]["Address"]).(string)
		item.RMobile = (maps[0]["RMobile"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertPacketTableItem(item *ZldPacketData) {
	// first, try to create the table
	CreateZldPacketTable()

	if len(item.ExpressNo) == 0 || !AlreadyHavePacketItem(item.ExpressNo) {
		// no express number
		DoInsertPacketTableItem(item)
	} 

	// No same sn item, do insert
	// if !AlreadyHavePacketItem(item.ExpressNo) {
	// 	DoInsertPacketTableItem(item)
	// }
}

func UpdatePacketTableItem(item *ZldPacketData) error{
	s := fmt.Sprintf("UPDATE `%s`", ZLD_PACKET_TBL_NAME)
	s = fmt.Sprintf("%s SET ", s)
	s = fmt.Sprintf("%s `TaskId` = '%s'", s, item.TaskId)
	s = fmt.Sprintf("%s, `Sender` = '%s'", s, item.SendTime)
	s = fmt.Sprintf("%s, `Receiver` = '%s'", s, item.Receiver)
	s = fmt.Sprintf("%s, `Address` = '%s'", s, item.Address)
	s = fmt.Sprintf("%s, `ExpressNo` = '%s'", s, item.ExpressNo)
	s = fmt.Sprintf("%s, `RMobile` = '%s'", s, item.RMobile)
	s = fmt.Sprintf("%s, `SendTime` = '%v'", s, item.SendTime)
	s = fmt.Sprintf("%s, `ExpressFee` = '%v'", s, item.ExpressFee)
	s = fmt.Sprintf("%s, `Comment` = '%s'", s, item.Comment)
	s = fmt.Sprintf("%s WHERE (`Id` = '%v');", s, item.Id)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		zllogs.WriteDebugLog("Update record(Id=%v) in table %s  ...... DONE", item.Id, ZLD_PACKET_TBL_NAME)
	} else {
		zllogs.WriteErrorLog("Update record(Id=%v) in table %s  ...... ERROR", item.Id, ZLD_PACKET_TBL_NAME)
	}

	return err
}

func SelectPacketTableItemWithTaskId(item *ZldPacketData) error{
	lenExpressNo := len(item.ExpressNo) 
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PACKET_TBL_NAME)
	if lenExpressNo == 0 && item.Id == -1 {
		// only task id
		s = fmt.Sprintf("%s WHERE (`TaskId` = '%s' AND `ExpressNo` = '');", s, item.TaskId)
	} else if lenExpressNo != 0 {
		s = fmt.Sprintf("%s WHERE (`ExpressNo` = '%s');", s, item.ExpressNo)
	} else if lenExpressNo == 0 && item.Id != -1 {
		s = fmt.Sprintf("%s WHERE (`TaskId` = '%s' AND `Id` = '%d');", s, item.TaskId, item.Id)
	}	
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		_packetId := (maps[0]["Id"]).(string)
		packetId, _ := strconv.ParseInt(_packetId, 10, 64)
		item.Id = int32(packetId)
		stime := (maps[0]["SendTime"]).(string)
		item.SendTime, _ = strconv.ParseInt(stime, 10, 64)
		fee := (maps[0]["ExpressFee"]).(string)
		item.ExpressFee, _ = strconv.ParseFloat(fee, 64)

		item.Sender = (maps[0]["Sender"]).(string)
		item.Receiver = (maps[0]["Receiver"]).(string)
		item.Address = (maps[0]["Address"]).(string)
		item.RMobile = (maps[0]["RMobile"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func AddPacketJSONString(json *simplejson.Json) {
	item := NewZldPacketDBData()

	item.TaskId, _ = json.Get("TaskId").String()
	item.Sender, _ = json.Get("Sender").String()
	item.Receiver, _ = json.Get("Receiver").String()
	item.Address, _ = json.Get("Address").String()
	item.ExpressNo, _ = json.Get("ExpressNo").String()
	item.RMobile, _ = json.Get("RMobile").String()
	item.SendTime, _ = json.Get("SendTime").Int64()
	item.ExpressFee, _ = json.Get("ExpressFee").Float64()
	item.Comment, _ = json.Get("Comment").String() 
	InsertPacketTableItem(item)	
}

func UpdatePacketJSONString(json *simplejson.Json) error {
	item := NewZldPacketDBData()

	packetId, _ := json.Get("Id").Int()
	item.Id = int32(packetId)
	item.TaskId, _ = json.Get("TaskId").String()
	item.Sender, _ = json.Get("Sender").String()
	item.Receiver, _ = json.Get("Receiver").String()
	item.Address, _ = json.Get("Address").String()
	item.ExpressNo, _ = json.Get("ExpressNo").String()
	item.RMobile, _ = json.Get("RMobile").String()
	item.SendTime, _ = json.Get("SendTime").Int64()
	item.ExpressFee, _ = json.Get("ExpressFee").Float64()
	item.Comment, _ = json.Get("Comment").String() 

	return UpdatePacketTableItem(item)
}

