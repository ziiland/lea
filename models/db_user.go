package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lea/zllogs"
	"strconv"
	"time"
)

const ZLD_USER_TBL_NAME string = "zld_user"

type ZldUserData struct {
	Id				int32 		`orm:"pk;auto"`
	UserId			string		// user id
	Password		string		// password
	Name			string		// 
	RealName		string		//
	NickName		string		//
	Mobile			string
	Mobile2			string
	Mailbox			string
	Mailbox2		string
	WechatId		string
	Address			string
	CellIds         string

	//Sex				string
	//IdentifyNo		string		//

	Type 			int64
	Rank			int64
	Fortune			int64
	CreateTime		int64
	TotalInvitors	int64
	CurrentInvitors	int64
	Invitee			string
	Invitors		string
	Signature		string
	Introduce		string
	Comment   		string
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func NewZldUserDBData() *ZldUserData{
	return &ZldUserData{}
}

func CreateZldUserTable() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, ZLD_USER_TBL_NAME)
	s = fmt.Sprintf("%s ( `Id` int(10) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `UserId` varchar(10) NOT NULL DEFAULT '' COMMENT '用户Id',", s)
	s = fmt.Sprintf("%s `Password` varchar(16) NOT NULL DEFAULT '' COMMENT '密码',", s)
	s = fmt.Sprintf("%s `Name` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',", s)
	s = fmt.Sprintf("%s `RealName` varchar(128) NOT NULL DEFAULT '' COMMENT '真实姓名',", s)
	s = fmt.Sprintf("%s `NickName` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',", s)
	s = fmt.Sprintf("%s `Mobile` varchar(16) NOT NULL DEFAULT '' COMMENT '手机',", s)
	s = fmt.Sprintf("%s `Mobile2` varchar(16) NOT NULL DEFAULT '' COMMENT '备用手机',", s)
	s = fmt.Sprintf("%s `Mailbox` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',", s)
	s = fmt.Sprintf("%s `Mailbox2` varchar(64) NOT NULL DEFAULT '' COMMENT '备用邮箱',", s)
	s = fmt.Sprintf("%s `WechatId` varchar(32) NOT NULL DEFAULT '' COMMENT '微信号',", s)
	s = fmt.Sprintf("%s `Address` varchar(256) NOT NULL DEFAULT '' COMMENT '地址',", s)
	s = fmt.Sprintf("%s `CellIds` text NOT NULL DEFAULT '' COMMENT '土地',", s)

	//s = fmt.Sprintf("%s `Sex` varchar(8) NOT NULL DEFAULT '' COMMENT '性别',", s)
	//s = fmt.Sprintf("%s `IdentifyNo` varchar(32) NOT NULL DEFAULT '' COMMENT '身份证号',", s)
	s = fmt.Sprintf("%s `Type` int(4) NOT NULL DEFAULT 0 COMMENT '种类',", s)
	s = fmt.Sprintf("%s `Rank` int(4) NOT NULL DEFAULT 0 COMMENT '等级',", s)
	s = fmt.Sprintf("%s `Fortune` int(10) NOT NULL DEFAULT 0 COMMENT '财富',", s)	
	s = fmt.Sprintf("%s `TotalInvitors` int(4) NOT NULL DEFAULT 0 COMMENT '总邀请数',", s)
	s = fmt.Sprintf("%s `CurrentInvitors` int(4) NOT NULL DEFAULT 0 COMMENT '当前可用邀请数',", s)
	s = fmt.Sprintf("%s `CreateTime` int(10) NOT NULL DEFAULT 0 COMMENT '加入时间',", s)

	s = fmt.Sprintf("%s `Invitee` text NOT NULL DEFAULT '' COMMENT '邀请人',", s)
	s = fmt.Sprintf("%s `Invitors` text NOT NULL DEFAULT '' COMMENT '被邀请人',", s)
	s = fmt.Sprintf("%s `Signature` text NOT NULL DEFAULT '' COMMENT '签名',", s)
	s = fmt.Sprintf("%s `Introduce` text NOT NULL DEFAULT '' COMMENT '介绍',", s)
	s = fmt.Sprintf("%s `Comment` text NOT NULL DEFAULT '' COMMENT '备注'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='用户表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", ZLD_USER_TBL_NAME)
		zllogs.WriteDebugLog("Create %s table ...... DONE", ZLD_USER_TBL_NAME)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		zllogs.WriteErrorLog("Create %s table ...... ERROR", ZLD_USER_TBL_NAME)
	}	
}

func SelectUserTableItem(item *ZldUserData) error{
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_USER_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`UserId` = '%s');", s, item.UserId)
	//fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	//fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		stime := (maps[0]["CreateTime"]).(string)
		item.CreateTime, _ = strconv.ParseInt(stime, 10, 64)
		fortune := (maps[0]["Fortune"]).(string)
		item.Fortune, _ = strconv.ParseInt(fortune, 10, 64)

		rank := (maps[0]["Rank"]).(string)
		item.Rank, _ = strconv.ParseInt(rank, 10, 64)
		invitors := (maps[0]["TotalInvitors"]).(string)
		item.TotalInvitors, _ = strconv.ParseInt(invitors, 10, 64)		
		cinvitors := (maps[0]["CurrentInvitors"]).(string)
		item.CurrentInvitors, _ = strconv.ParseInt(cinvitors, 10, 64)	
		ctype := (maps[0]["Type"]).(string)
		item.Type, _ = strconv.ParseInt(ctype, 10, 64)	

		item.Password = (maps[0]["Password"]).(string)
		item.Name = (maps[0]["Name"]).(string)
		item.RealName = (maps[0]["RealName"]).(string)
		item.NickName = (maps[0]["NickName"]).(string)
		item.Mobile = (maps[0]["Mobile"]).(string)
		item.Mobile2 = (maps[0]["Mobile2"]).(string)
		item.Mailbox = (maps[0]["Mailbox"]).(string)
		item.Mailbox2 = (maps[0]["Mailbox2"]).(string)
		item.WechatId = (maps[0]["WechatId"]).(string)	
		item.Address = (maps[0]["Address"]).(string)
		item.CellIds = (maps[0]["CellIds"]).(string)
		fmt.Println("SelectItem=", *item)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err	
}

func AlreadyHaveUserItem(id string) bool {
	s := fmt.Sprintf("SELECT * FROM `%s`", ZLD_USER_TBL_NAME)
	s = fmt.Sprintf("%s WHERE (`UserId` = '%s' );", s, id)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)

	if err == nil && num > 0 {
		return true
	} else {
		return false
	}
}

func DoInsertUserTableItem(item *ZldUserData) {
	// insert data to the table
	s := fmt.Sprintf("INSERT INTO `%s`", ZLD_USER_TBL_NAME)
	s = fmt.Sprintf("%s (`UserId`, `Password`, `Name`, `RealName`, `NickName`, `Mobile`", s)
	s = fmt.Sprintf("%s, `Mobile2`, `Mailbox`, `Mailbox2`, `WechatId`, `Address`, `CellIds`", s)
	s = fmt.Sprintf("%s, `Type`, `Rank`, `Fortune`, `CreateTime`, `TotalInvitors`, `CurrentInvitors`", s)
	s = fmt.Sprintf("%s, `Invitee`, `Invitors`, `Signature`, `Introduce`, `Comment`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s'", s, item.UserId, item.Password, item.Name, item.RealName, item.NickName, item.Mobile)
	s = fmt.Sprintf("%s , '%s', '%s', '%s', '%s', '%s', '%s'", s, item.Mobile2, item.Mailbox, item.Mailbox2, item.WechatId, item.Address, item.CellIds)
	s = fmt.Sprintf("%s , '%v', '%v', '%v', '%v', '%v', '%v'", s, item.Type, item.Rank, item.Fortune, time.Now().Unix(), item.TotalInvitors, item.CurrentInvitors)
	s = fmt.Sprintf("%s , '%s', '%s', '%s', '%s', '%s');", s, item.Invitee, item.Invitors, item.Signature, item.Introduce, item.Comment)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		zllogs.WriteDebugLog("Insert a record to %s table ...... DONE", ZLD_USER_TBL_NAME)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		zllogs.WriteErrorLog("Insert a record to %s table ...... ERROR", ZLD_USER_TBL_NAME)
	}		
}

func InsertUserTableItem(item *ZldUserData) {
	// first, try to create the table
	CreateZldUserTable()

	// No same sn item, do insert
	if !AlreadyHaveUserItem(item.UserId) {
		DoInsertUserTableItem(item)
	}
}
