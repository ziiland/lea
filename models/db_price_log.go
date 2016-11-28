package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/common"
	"lea/zllogs"
	"strconv"
	"time"
)

const ZLD_PRICE_LOG_TBL_NAME string = "zld_price_log"

type ZldPriceLogData struct {
	Id				int32 		`orm:"pk;auto"`
	Kind			string		// 
	Action			string
	OperatorId		string
	ActionTime		int64
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldPriceLogDBData() *ZldPriceLogData{
	return &ZldPriceLogData{}
}

func CreateZldPriceLogTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_PRICE_LOG_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `Kind` varchar(16) NOT NULL DEFAULT '' COMMENT '品种',", s)
	s = fmt.Sprintf("%s `Action` varchar(16) NOT NULL DEFAULT '' COMMENT '动作',", s)
	s = fmt.Sprintf("%s `OperatorId` varchar(32) NOT NULL DEFAULT '' COMMENT '工作员号',", s)
	s = fmt.Sprintf("%s `ActionTime` int(10) NOT NULL DEFAULT 0 COMMENT '动作时间',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='价格日志表';", s)
	fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_PRICE_LOG_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_PRICE_LOG_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_PRICE_LOG_TBL_NAME)
	}	
}

func DoInsertPriceLogTableItem(item *ZldPriceLogData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_PRICE_LOG_TBL_NAME)
	s = fmt.Sprintf("%s (`Kind`, `Action`, `OperatorId`, `ActionTime`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ('%s', '%s', '%s'", s, item.Kind, item.Action, item.OperatorId)
	if item.ActionTime == 0 {
		s = fmt.Sprintf("%s, '%v'", s, time.Now().Unix())
	} else {
		s = fmt.Sprintf("%s, '%v'", s, item.ActionTime)
	}
	s = fmt.Sprintf("%s , '%s');", s, item.Comment)
	fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_PRICE_LOG_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_PRICE_LOG_TBL_NAME)
	}		
}

func AlreadyHavePriceLogItem(kind string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PRICE_LOG_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`Kind` = '%s' );", s, kind)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func InsertPriceLogTableItem(item *ZldPriceLogData) {
	// first, try to create the table
	CreateZldPriceLogTable()
	DoInsertPriceLogTableItem(item)
}

func DecodePriceLogOrmParamsToData(para orm.Params) (item ZldPriceLogData){
	item.Kind = (para["Kind"]).(string)
	id := (para["Id"]).(string)
	nId, _ := strconv.Atoi(id)
	item.Id = int32(nId)

	ctime := (para["ActionTime"]).(string)
	item.ActionTime, _ = strconv.ParseInt(ctime, 10, 64)

	item.Action = (para["Action"]).(string)
	item.OperatorId = (para["OperatorId"]).(string)
	item.Comment = (para["Comment"]).(string)

	return		
}

func QueryMatchPriceLogItemNums(kind string)(int64, error) {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PRICE_LOG_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`Kind` = '%s');", s, kind)	
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	return o.Raw(s).Values(&maps);
}

func SelectPriceLogTableItemsWithKind(kind string) ([]ZldPriceLogData, error){
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PRICE_LOG_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`Kind` = '%s' );", s, kind)
	fmt.Println("s=", s)

	var maps []orm.Params
	var num int64
	var err error
	o := orm.NewOrm()
	num, err = o.Raw(s).Values(&maps);

	logs := make([]ZldPriceLogData, num) 
	if err == nil && num > 0 {
		for i, v := range maps {
			logs[i] = DecodePriceLogOrmParamsToData(v)
		}
	}

	fmt.Println("logs=", logs)
	return logs, err	
}

func HandleStandardPriceLogItem(kind, command, worker string) {
	item := NewZldPriceLogDBData()

	item.Kind = kind
	item.OperatorId = worker
	switch command {
	case common.ZLD_CMD_ADD_RRICE:
		item.Action = common.ZLD_ACTION_ADD
	case common.ZLD_CMD_UPDATE_PRICE:
		item.Action = common.ZLD_ACTION_UPDATE
	case common.ZLD_CMD_DEL_PRICE:
		item.Action = common.ZLD_ACTION_DEL
	}

	InsertPriceLogTableItem(item)	
}