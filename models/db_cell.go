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
	OwnerId			string
	CellId			string
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
	s = fmt.Sprintf("%s `OwnerId` varchar(10) NOT NULL DEFAULT '' COMMENT '所有人',", s)
	s = fmt.Sprintf("%s `CellId` varchar(8) NOT NULL DEFAULT '' COMMENT '单元号',", s)
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
	s = fmt.Sprintf("%s (`FarmId`, `OwnerId`, `CellId`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s');", s, item.FarmId, item.OwnerId, item.CellId, item.Comment)
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

