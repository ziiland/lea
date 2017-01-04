package common

// import (
// 	"github.com/astaxie/beego"
// )

const ZLD_PATH_LOGIN string  = "/land/login"
const ZLD_PATH_WORKER string = "/land/worker"
const ZLD_PATH_TASK string = "/land/task"
const ZLD_PATH_PRICE string = "/land/price"
const ZLD_PATH_CELL string = "/land/cell"
const ZLD_PATH_PACKET string = "/land/packet"

const ZLD_CMD_LOGIN string = "Login"
const ZLD_CMD_LOAD_PARA string = "LoadPara"
const ZLD_CMD_LOAD_VER string = "LoadVer"
const ZLD_CMD_UNLOAD string = "UnLoad"
///////////////////////////////////////////////////////////////////////////////
const ZLD_CMD_LOAD_PRICE string = "LoadPrice"
const ZLD_CMD_ADD_RRICE string = "AddPrice"
const ZLD_CMD_UPDATE_PRICE string = "UpdatePrice"
const ZLD_CMD_DEL_PRICE string = "DelPrice"
///////////////////////////////////////////////////////////////////////////////
const ZLD_CMD_LOAD_WORKER string = "LoadWorker"
const ZLD_CMD_ADD_WORKER string = "AddWorker"
const ZLD_CMD_DEL_WORKER string = "DelWorker"
const ZLD_CMD_UPD_WORKER string = "UpdateWorker"
const ZLD_CMD_CHGPWD_WORKER string = "ChgPwd"
///////////////////////////////////////////////////////////////////////////////
const ZLD_CMD_LOAD_TASK string = "LoadTask"
const ZLD_CMD_ADD_TASK string = "AddTask"
const ZLD_CMD_ASSIGN_TASK string = "AssignTask"
const ZLD_CMD_SUBMIT_TASK string = "SubmitTask"
const ZLD_CMD_CHECK_TASK string = "CheckTask"
const ZLD_CMD_CLOSE_TASK string = "CloseTask"
const ZLD_CMD_TRANS_TASK string = "TransferTask"
const ZLD_CMD_QUERY_TASK string = "QueryTask"
const ZLD_CMD_ARCHIVE_TASK string = "ArchiveTask"
const ZLD_CMD_CANCEL_TASK string = "CancelTask"
const ZLD_CMD_BEGIN_TASK string = "BeginTask"
///////////////////////////////////////////////////////////////////////////////
const ZLD_CMD_CELL_BINDNFC string = "BindNFC"
///////////////////////////////////////////////////////////////////////////////
const ZLD_CMD_ADD_PACKET string = "AddPacket"
const ZLD_CMD_GET_PACKET string = "GetPacket"
const ZLD_CMD_SET_PACKET string = "SetPacket"

const ZLD_PARA_COMMAND string = "Command"
const ZLD_PARA_WORKER string = "Worker"
const ZLD_PARA_CHECKER string = "Checker"
const ZLD_PARA_PWD string = "Password"
const ZLD_PARA_TITLE string = "Title"
const ZLD_PARA_FARM string = "Farm"
const ZLD_PARA_CELL string = "Cell"
const ZLD_PARA_PATCH string = "Patch"
const ZLD_PARA_NAME string = "Name"
const ZLD_PARA_SEX string = "Sex"
const ZLD_PARA_ID string = "IdentifyNo"
const ZLD_PARA_COMMENT string = "Comment"
const ZLD_PARA_TASKID string = "TaskId"
const ZLD_PARA_STIME string = "STime"
const ZLD_PARA_ETIME string = "ETime"
const ZLD_PARA_STATE string = "State"
const ZLD_PARA_TYPE string = "Type"
const ZLD_PARA_KIND string = "Kind"
const ZLD_PARA_SHOW string = "Show"
const ZLD_PARA_PRICE string = "Price"
const ZLD_PARA_DISCOUNT string = "Discount"
const ZLD_PARA_NFC string = "NFC"
const ZLD_PARA_EXPNO string = "ExpressNo"
const ZLD_PARA_PACKEDID string = "PacketId"

const ZLD_STR_ON string = "on"
const ZLD_STR_OFF string = "off"
const ZLD_STR_OK string = "ok"
const ZLD_STR_ADMIN string = "Admin"
const ZLD_STR_MANAGER string = "Manager"
const ZLD_STR_WORKER string = "Worker"

const ZLD_ACTION_ADD string = "Add"
const ZLD_ACTION_UPDATE string = "Update"
const ZLD_ACTION_DEL string = "Delete"

// task action
const ZLD_TASK_ACTION_ADD string = "Add"
const ZLD_TASK_ACTION_LIST string = "List"
const ZLD_TASK_ACTION_SEARCH string = "Search"
const ZLD_TASK_ACTION_DETAIL string = "Detail"
const ZLD_TASK_ACTION_START string = "Start"
const ZLD_TASK_ACTION_SUBMIT string = "Submit"
const ZLD_TASK_ACTION_CHECK string = "Check"
const ZLD_TASK_ACTION_ASSIGN string = "Assign"
const ZLD_TASK_ACTION_CLOSE string = "Close"
const ZLD_TASK_ACTION_CANCEL string = "Cancel"
const ZLD_TASK_ACTION_ARCHIVE string = "Archive" 
const ZLD_TASK_ACTION_NOTIFY string = "Notify"

// task states
const ZLD_TASK_STATE_CREATED string = "Created"
const ZLD_TASK_STATE_ASSIGNED string = "Assigned"
const ZLD_TASK_STATE_STARTED string = "Started"
const ZLD_TASK_STATE_FINISHED string = "Finished"
const ZLD_TASK_STATE_CHECKED string = "Checked"
const ZLD_TASK_STATE_CLOSED string = "Closed"
const ZLD_TASK_STATE_CANCELED string = "Canceled"
const ZLD_TASK_STATE_ARCHIVED string = "Archived"

// task type
const ZLD_TASK_TYPE_SOIL = 0
const ZLD_TASK_TYPE_SEED = 1
const ZLD_TASK_TYPE_WATER = 2
const ZLD_TASK_TYPE_FERTI = 3
const ZLD_TASK_TYPE_SCAFFOLD = 4
const ZLD_TASK_TYPE_TRANS = 5
const ZLD_TASK_TYPE_GRAFTING = 6
const ZLD_TASK_TYPE_WEEDING = 7
const ZLD_TASK_TYPE_DEINSECT = 8
const ZLD_TASK_TYPE_HARVEST = 9
const ZLD_TASK_TYPE_EXPRESS = 10

var ZldTaskType [11]string = [11]string{"翻地", "播种", "浇水", "施肥", "搭架子", "移栽", "嫁接", "除草", "除虫", "收割", "快递"}




