package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	//"strconv"
	//"time"
)

const ZLD_NOTIFY_TBL_NAME string = "zld_notify"

type ZldNotifyData struct {
	Id				int32 		`orm:"pk;auto"`
	Area			string	
	CreatorId 		string
	CreateTime		int64	
	Content			string 
	Status			bool    
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldNotifyDBData() *ZldNotifyData{
	return &ZldNotifyData{}
}

func CreateZldNotifyTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_NOTIFY_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `Area` varchar(32) NOT NULL DEFAULT '' COMMENT '产地',", s)
	s = fmt.Sprintf("%s `CreatorId` varchar(10) NOT NULL DEFAULT '' COMMENT '创建者',", s)
	s = fmt.Sprintf("%s `Status` bool NOT NULL DEFAULT 0.0 COMMENT '发送',", s)
	s = fmt.Sprintf("%s `CreateTime` int(10) NOT NULL DEFAULT 0 COMMENT '创建时间',", s)
	s = fmt.Sprintf("%s `Content` text NOT NULL DEFAULT '' COMMENT '内容',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='消息表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_NOTIFY_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_NOTIFY_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_GOODS_TBL_NAME)
	}	
}

func DoInsertNotifyTableItem(item *ZldNotifyData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_NOTIFY_TBL_NAME)
	s = fmt.Sprintf("%s (`Area`, `CreatorId`, `Status`, `CreateTime`, `Content`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%v', '%v', '%s', '%s');", s, item.Area, item.CreatorId, item.Status, item.CreateTime, item.Content, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_NOTIFY_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_NOTIFY_TBL_NAME)
	}		
}

// func AlreadyHavePacketItem(ExpressNo string) bool {
// 	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_GOODS_TBL_NAME)
// 	s = fmt.Sprintf("%s WHERE (`ExpressNo` = '%s');", s, ExpressNo)

// 	var maps []orm.Params
// 	o := orm.NewOrm()
// 	num, err := o.Raw(s).Values(&maps)

// 	if err == nil && num > 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func SelectPacketTableItem(item *ZldNotifyData) error{
// 	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_GOODS_TBL_NAME)
// 	s = fmt.Sprintf("%s WHERE (`ExpressNo` = '%s');", s, item.ExpressNo)
// 	//fmt.Println("s=", s)

// 	var maps []orm.Params
// 	o := orm.NewOrm()
// 	num, err := o.Raw(s).Values(&maps)
// 	//fmt.Printf("num=%d, maps=%v\n", num, maps)

// 	if err == nil && num > 0 {
// 		stime := (maps[0]["SendTime"]).(string)
// 		item.SendTime, _ = strconv.ParseInt(stime, 10, 64)
// 		fee := (maps[0]["ExpressFee"]).(string)
// 		item.ExpressFee, _ = strconv.ParseFloat(fee, 64)

// 		item.Sender = (maps[0]["Sender"]).(string)
// 		item.Receiver = (maps[0]["Receiver"]).(string)
// 		item.Address = (maps[0]["Address"]).(string)
// 		item.Comment = (maps[0]["Comment"]).(string)	
// 		fmt.Println("SelectItem=", *item)
// 	} else {
// 		fmt.Println("Select NONE! Error=%v", err)
// 	}

// 	return err	
// }

// func InsertPacketTableItem(item *ZldNotifyData) {
// 	// first, try to create the table
// 	CreateZldPriceTable()

// 	// No same sn item, do insert
// 	if !AlreadyHavePacketItem(item.ExpressNo) {
// 		DoInsertPacketTableItem(item)
// 	}
// }

