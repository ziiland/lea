package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	//"strconv"
	//"time"
)

const ZLD_GOODS_TBL_NAME string = "zld_goods"

type ZldGoodsData struct {
	Id				int32 		`orm:"pk;auto"`
	Area			string		
	OwnerId			string      
	Name			string
	Weight			float64
	HarvestTime		int64
	Packets			string    // serialized packets
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldGoodsDBData() *ZldGoodsData{
	return &ZldGoodsData{}
}

func CreateZldGoodsTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_GOODS_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `Area` varchar(32) NOT NULL DEFAULT '' COMMENT '产地',", s)
	s = fmt.Sprintf("%s `OwnerId` varchar(10) NOT NULL DEFAULT '' COMMENT '所有人',", s)
	s = fmt.Sprintf("%s `Name` varchar(64) NOT NULL DEFAULT '' COMMENT '货物',", s)
	s = fmt.Sprintf("%s `Weight` float NOT NULL DEFAULT 0.0 COMMENT '重量',", s)
	s = fmt.Sprintf("%s `HarvestTime` int(10) NOT NULL DEFAULT 0 COMMENT '收割时间',", s)
	s = fmt.Sprintf("%s `Packets` text NOT NULL DEFAULT '' COMMENT '包裹',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='收获表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_GOODS_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_GOODS_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_GOODS_TBL_NAME)
	}	
}

func DoInsertGoodsTableItem(item *ZldGoodsData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_GOODS_TBL_NAME)
	s = fmt.Sprintf("%s (`Area`, `OwnerId`, `Name`, `Weight`, `HarvestTime`, `Packets`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%v', '%v', '%s', '%s');", s, item.Area, item.OwnerId, item.Name, item.Weight, item.HarvestTime, item.Packets, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_GOODS_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_GOODS_TBL_NAME)
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

// func SelectPacketTableItem(item *ZldGoodsData) error{
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

// func InsertPacketTableItem(item *ZldGoodsData) {
// 	// first, try to create the table
// 	CreateZldPriceTable()

// 	// No same sn item, do insert
// 	if !AlreadyHavePacketItem(item.ExpressNo) {
// 		DoInsertPacketTableItem(item)
// 	}
// }

