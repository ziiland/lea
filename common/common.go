package common

// import (
// 	"github.com/astaxie/beego"
// )

const ZLD_PATH_LOGIN string  = "/land/login"
const ZLD_PATH_WORKER string = "/land/worker"
const ZLD_PATH_TASK string = "/land/task"
const ZLD_PATH_PRICE string = "/land/price"

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
const ZLD_CMD_ADD_WORKER = "AddWorker"
const ZLD_CMD_DEL_WORKER = "DelWorker"
const ZLD_CMD_UPD_WORKER = "UpdateWorker"
const ZLD_CMD_CHGPWD_WORKER = "ChgPwd"
///////////////////////////////////////////////////////////////////////////////
const ZLD_CMD_LOAD_TASK string = "LoadTask"
const ZLD_CMD_ADD_TASK = "AddTask"
const ZLD_CMD_ASSIGN_TASK = "AssignTask"
const ZLD_CMD_SUBMIT_TASK= "SubmitTask"
const ZLD_CMD_CHECK_TASK = "CheckTask"
const ZLD_CMD_CLOSE_TASK = "CloseTask"
const ZLD_CMD_TRANS_TASK = "TransferTask"
const ZLD_CMD_QUERY_TASK = "QueryTask"
const ZLD_CMD_ARCHIVE_TASK = "ArchiveTask"
const ZLD_CMD_CANCEL_TASK = "CancelTask"
const ZLD_CMD_BEGIN_TASK = "BeginTask"
///////////////////////////////////////////////////////////////////////////////

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

const ZLD_STR_ON string = "on"
const ZLD_STR_OFF string = "off"
const ZLD_STR_OK string = "ok"
const ZLD_STR_ADMIN string = "Admin"
const ZLD_STR_MANAGER string = "Manager"
const ZLD_STR_WORKER string = "Worker"

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




