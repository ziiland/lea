package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	"strconv"
	//"time"
)

const ZLD_PRICE_TBL_NAME string = "zld_price"

type ZldPriceData struct {
	Id				int32 		`orm:"pk;auto"`
	Name			string		// goods/service name
	Kind			string      // goods/service type
	Price			float64
	Discount		float64
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldPriceDBData() *ZldPriceData{
	return &ZldPriceData{}
}

func CreateZldPriceTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_PRICE_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `Name` varchar(128) NOT NULL DEFAULT '' COMMENT '名称',", s)
	s = fmt.Sprintf("%s `Kind` varchar(32) NOT NULL DEFAULT '' COMMENT '种类',", s)
	s = fmt.Sprintf("%s `Price` float NOT NULL DEFAULT 0.0 COMMENT '价格',", s)
	s = fmt.Sprintf("%s `Discount` float NOT NULL DEFAULT 0.0 COMMENT '折扣',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='价格表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_PRICE_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_PRICE_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_PRICE_TBL_NAME)
	}	
}

func DoInsertPriceTableItem(item *ZldPriceData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_PRICE_TBL_NAME)
	s = fmt.Sprintf("%s (`Name`, `Kind`, `Price`, `Discount`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%v', '%v', '%s');", s, item.Name, item.Kind, item.Price, item.Discount, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_PRICE_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_PRICE_TBL_NAME)
	}		
}

func AlreadyHavePriceItem(kind string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PRICE_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`Kind` = '%s');", s, kind)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func SelectPriceTableItem(item *ZldPriceData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_PRICE_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`Kind` = '%s');", s, item.Kind)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		price := (maps[0]["Price"]).(string)
		item.Price, _ = strconv.ParseFloat(price, 64)
		discount := (maps[0]["Discount"]).(string)
		item.Discount, _ = strconv.ParseFloat(discount, 64)

		item.Name = (maps[0]["Name"]).(string)
		item.Comment = (maps[0]["Comment"]).(string)	
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func InsertPriceTableItem(item *ZldPriceData) {
	// first, try to create the table
	CreateZldPriceTable()

	// No same sn item, do insert
	if !AlreadyHavePriceItem(item.Kind) {
		DoInsertPriceTableItem(item)
	}
}

func DecodePriceOrmParamsToData(para orm.Params) (item ZldPriceData){
	id := (para["Id"]).(string)
	nId, _ := strconv.Atoi(id)
	item.Id = int32(nId)

	cprice := (para["Price"]).(string)
	item.Price, _ = strconv.ParseFloat(cprice, 64)
	discount := (para["Discount"]).(string)
	item.Discount, _ = strconv.ParseFloat(discount, 64)

	item.Name = (para["Name"]).(string)
	item.Kind = (para["Kind"]).(string)
	item.Comment = (para["Comment"]).(string)

	return		
}

func QueryAllPriceTableItem() ([]ZldPriceData, error) {
	s := fmt.Sprintf("SELECT * FROM `%s`;", ZLD_PRICE_TBL_NAME)
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	prices := make([]ZldPriceData, num)
	if err == nil && num > 0 {
		for i, v := range maps {
			prices[i] = DecodePriceOrmParamsToData(v)
		}
	}
	fmt.Println("prices=", prices)
	return prices, err
}
