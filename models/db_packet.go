package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	"strconv"
	//"time"
)

const ZLD_PACKET_TBL_NAME string = "zld_packet"

type ZldPacketData struct {
	Id				int32 		`orm:"pk;auto"`
	Sender			string		
	Receiver		string      
	Address			string
	ExpressNo		string
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
	s = fmt.Sprintf("%s `Sender` varchar(128) NOT NULL DEFAULT '' COMMENT '发件人',", s)
	s = fmt.Sprintf("%s `Receiver` varchar(128) NOT NULL DEFAULT '' COMMENT '收件人',", s)
	s = fmt.Sprintf("%s `Address` varchar(256) NOT NULL DEFAULT '' COMMENT '地址',", s)
	s = fmt.Sprintf("%s `ExpressNo` varchar(32) NOT NULL DEFAULT '' COMMENT '快递单号',", s)
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
	s = fmt.Sprintf("%s (`Sender`, `Receiver`, `Address`, `ExpressNo`, `SendTime`, `ExpressFee`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%v', '%v', '%s');", s, item.Sender, item.Receiver, item.Address, item.ExpressNo, item.SendTime, item.ExpressFee, item.Comment)
	//fmt.Println("s=", s)

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
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertPacketTableItem(item *ZldPacketData) {
	// first, try to create the table
	CreateZldPriceTable()

	// No same sn item, do insert
	if !AlreadyHavePacketItem(item.ExpressNo) {
		DoInsertPacketTableItem(item)
	}
}

