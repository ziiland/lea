package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	//"strconv"
	//"time"
)

const ZLD_PATCH_TBL_NAME string = "zld_patch"

type ZldPatchData struct {
	Id				int32 		`orm:"pk;auto"`
	FarmId			string		// farm id
	CellId			string
	PatchId			string
	SOP				string
	State 			string
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldPatchDBData() *ZldPatchData{
	return &ZldPatchData{}
}

func CreateZldPatchTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_PATCH_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `FarmId` varchar(6) NOT NULL DEFAULT '' COMMENT '农场号',", s)
	s = fmt.Sprintf("%s `CellId` varchar(8) NOT NULL DEFAULT '' COMMENT '单元号',", s)
	s = fmt.Sprintf("%s `PatchId` varchar(1) NOT NULL DEFAULT '' COMMENT '小片号',", s)
	s = fmt.Sprintf("%s `SOP` varchar(128) NOT NULL DEFAULT '' COMMENT 'SOP',", s)
	s = fmt.Sprintf("%s `State` varchar(32) NOT NULL DEFAULT '' COMMENT '状态',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='小片表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_PATCH_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_PATCH_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_PATCH_TBL_NAME)
	}	
}

func DoInsertPatchTableItem(item *ZldPatchData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_PATCH_TBL_NAME)
	s = fmt.Sprintf("%s (`FarmId`, `CellId`, `PatchId`, `SOP`, `State`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s');", s, item.FarmId, item.CellId, item.PatchId, item.SOP, item.State, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_PATCH_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_PATCH_TBL_NAME)
	}		
}

func AlreadyHavePatchItem(farm, cell, patch string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PATCH_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s' AND `CellId` = '%s' AND `PatchId` = '%s');", s, farm, cell, patch)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func SelectPatchTableItem(item *ZldPatchData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PATCH_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`FarmId` = '%s' AND `CellId` = '%s' AND `Patched` = '%s');", s, item.FarmId, item.CellId, item.PatchId)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		item.SOP = (maps[0]["SOP"]).(string)
		item.State = (maps[0]["State"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertPatchTableItem(item *ZldPatchData) {
	// first, try to create the table
	CreateZldPatchTable()

	// No same sn item, do insert
	if !AlreadyHavePatchItem(item.FarmId, item.CellId, item.PatchId) {
		DoInsertPatchTableItem(item)
	}
}

