package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	//"strconv"
	//"time"
)

const ZLD_CELL_TBL_NAME string = "zld_cell"

type ZldCellData struct {
	Id				int32 		`orm:"pk;auto"`
	FarmId			string		// farm id
	CellId			string
	NFCId			string
	OwnerId			string

	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldCellDBData() *ZldCellData{
	return &ZldCellData{}
}

func CreateZldCellTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_CELL_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `FarmId` varchar(6) NOT NULL DEFAULT '' COMMENT '农场号',", s)
	s = fmt.Sprintf("%s `CellId` varchar(8) NOT NULL DEFAULT '' COMMENT '单元号',", s)
	s = fmt.Sprintf("%s `NFCId` varchar(16) NOT NULL DEFAULT '' COMMENT '绑定NFC',", s)	
	s = fmt.Sprintf("%s `OwnerId` varchar(10) NOT NULL DEFAULT '' COMMENT '所有人',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='单元格表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_CELL_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_CELL_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_CELL_TBL_NAME)
	}	
}

func DoInsertCellTableItem(item *ZldCellData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_CELL_TBL_NAME)
	s = fmt.Sprintf("%s (`FarmId`, `CellId`, `NFCId`, `OwnerId`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s');", s, item.FarmId, item.CellId, item.NFCId, item.OwnerId, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_CELL_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_CELL_TBL_NAME)
	}		
}

func UpdateCellTableItem(item *ZldCellData) error {
	s := fmt.Sprintf("UPDATE `%s`", ZLD_CELL_TBL_NAME)
	s = fmt.Sprintf("%s SET ", s)
	s = fmt.Sprintf("%s, `NFCId` = '%s'", s, item.NFCId)
	s = fmt.Sprintf("%s, `OwnerId` = '%s'", s, item.OwnerId)	
	s = fmt.Sprintf("%s, `Comment` = '%s'", s, item.Comment)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s' AND `CellId` = '%s');", s, item.FarmId, item.CellId)
	fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		zllogs.WriteDebugLog("Update record(farm=%s, cell=%s) in table %s  ...... DONE", item.FarmId, item.CellId, ZLD_CELL_TBL_NAME)
	} else {
		zllogs.WriteErrorLog("Update record(farm=%s, cell=%s) in table %s  ...... ERROR", item.FarmId, item.CellId, ZLD_CELL_TBL_NAME)
	}

	return err
}


func AlreadyHaveCellItem(farm, cell string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_CELL_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s' AND `CellId` = '%s' );", s, farm, cell)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func SelectCellTableItem(item *ZldCellData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_CELL_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s' AND `CellId` = '%s');", s, item.FarmId, item.CellId)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		item.NFCId = (maps[0]["NFCId"]).(string)
		item.OwnerId = (maps[0]["OwnerId"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertCellTableItem(item *ZldCellData) {
	// first, try to create the table
	CreateZldCellTable()

	// No same sn item, do insert
	if !AlreadyHaveCellItem(item.FarmId, item.CellId) {
		DoInsertCellTableItem(item)
	}
}

func GetNFCContentInTableItem(farm, cell string) (nfc string) {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_CELL_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s' AND `CellId` = '%s');", s, farm, cell)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		nfc = (maps[0]["NFCId"]).(string)
		fmt.Println("nfc=", nfc)
	} else {
		nfc = ""
		fmt.Println("Select NONE! Error=%v", err)
	}
	return nfc		
}

// bind cell & nfc
func BindCellNFCTableItem(farm, cell, nfc string) int64 {
	var ret int64 = 0;   // 0 --- OK, 1 --- cell already have nfc

	CreateZldCellTable()

	item := NewZldCellDBData()
	item.FarmId = farm
	item.CellId = cell

	if !AlreadyHaveCellItem(item.FarmId, item.CellId) {
		// don't have cell, insert one item & bind nfc
		fmt.Println("Don't have the cell item, Insert One!")
		item.NFCId = nfc
		DoInsertCellTableItem(item)
	} else {
		fmt.Println("Have the cell item")
		SelectCellTableItem(item)
		if item.NFCId == "" {
			fmt.Println("No NFC record, Update the DataItem")
			item.NFCId = nfc
			UpdateCellTableItem(item)
		} else if (item.NFCId != nfc) {
			fmt.Println("Already Have NFC record, Conflict")
			ret = 1
		}
	}
	return ret;
}

