//// key constant
var KEY_ID = "Id";
var KEY_LOGIN = "Login";
var KEY_WORKER = "Worker";
var KEY_WORKERS = "Workers";
var KEY_WORKERID = "WorkerId";
var KEY_TITLE = "Title";
var KEY_NAME = "Name";
var KEY_FARM = "Farm";
var KEY_TASKS = "Tasks";
var KEY_COMMAND = "Command";
var KEY_PASSWORD = "Password";
var KEY_LOGS = "Logs";
var KEY_CHECKINTIME = "CheckInTime";
var KEY_CHECKOUTTIME = "CheckOutTime";
var KEY_ERRCODE = "Errcode";

var KEY_TASK_TASKID = "TaskId";
var KEY_TASK_SPONSORID = "SponsorId";
var KEY_TASK_FARMID = "FarmId";
var KEY_TASK_CELLID = "CellId";
var KEY_TASK_PATCHID = "PatchId";
var KEY_TASK_WORKERID = "WorkerId";
var KEY_TASK_CHECKERID = "CheckerId";
var KEY_TASK_STATE = "State";
var KEY_TASK_TYPE = "Type";
var KEY_TASK_CREATETIME = "CreateTime";
var KEY_TASK_STARTTIME = "StartTime";
var KEY_TASK_ENDTIME = "EndTime";
var KEY_TASK_CHECKTIME = "CheckTime";
var KEY_TASK_SCORE = "Score";
var KEY_TASK_UCOMMENT = "UserComment";
var KEY_TASK_COMMENT = "Comment";

var KEY_LOG_ACTION = "Action";
var KEY_LOG_OPERATORID = "OperatorId";
var KEY_LOG_ACTIONTIME = "ActionTime";

var KEY_PRICES = "Prices";

///////////////////////////////////////////////////////////////////////////////
// command
var CMD_LOAD_PARA = "LoadPara";
var CMD_LOAD_TASK = "LoadTask";
var CMD_LOAD_WORKER = "LoadWorker";
var CMD_UNLOAD = "UnLoad";
var CMD_LOGIN = "Login";

var CMD_ADD_WORKER = "AddWorker";
var CMD_DEL_WORKER = "DelWorker";
var CMD_UPD_WORKER = "UpdateWorker";
var CMD_CHGPWD_WORKER = "ChgPwd";

var CMD_QUERY_TASK = "QueryTask";
var CMD_ARCHIVE_TASK= "ArchiveTask";
var CMD_ADD_TASK = "AddTask";
var CMD_CANCEL_TASK = "CancelTask";
var CMD_ASSIGN_TASK = "AssignTask";
var CMD_CHECK_TASK = "CheckTask";
var CMD_CLOSE_TASK = "CloseTask";
var CMD_BEGIN_TASK = "BeginTask";
var CMD_SUBMIT_TASK = "SubmitTask";

var CMD_LOAD_PRICE = "LoadPrice";
var CMD_ADD_RRICE = "AddPrice";
var CMD_UPDATE_PRICE = "UpdatePrice";
var CMD_DEL_PRICE = "DelPrice";

///////////////////////////////////////////////////////////////////////////////
// url
var URL_LOGIN = "/land/login";
var URL_WORKER = "/land/worker";
var URL_TASK = "/land/task";
var URL_PRICE = "/land/price";

///////////////////////////////////////////////////////////////////////////////
// event
var EVT_HIDE_BTNADD = "HideBtnAdd";
var EVT_PARA_LOADED = "ParaLoaded";
var EVT_TASKS_LOADED = "TaskLoaded";

///////////////////////////////////////////////////////////////////////////////
// string
var STR_DEFAULT_PWD = "888888";
var STR_ADMIN = "Admin";
var STR_MANAGER = "Manager";
var STR_WORKER = "Worker";
var STR_ON = "on";
var STR_OFF = "off";

///////////////////////////////////////////////////////////////////////////////
var gTaskStateDes = {Created:"已创建", Assigned:"已分配", Started:"进行中", Finished:"已完成", Checked:"已检查", Closed:"已关闭", Canceled:"已取消", Archived:"已归档"};
var gTaskTypes = ["翻地", "播种", "浇水", "施肥", "搭架子", "移栽", "嫁接", "除草", "除虫", "收割", "快递"];
var gRoleDes = {"Admin":"管理员", "Manager":"经理", "Worker":"职员"};
var gLoginInfo = {workerId:"", title:""};

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
$(function() {
    var header = '<h1>Ziiland生产管理软件 <small>v1.0</small></h1>';
    var footer = '版权信息：寸田尺园网络科技(上海)有限公司';

    //显示页头和页尾
    $("#myheader").append(header);
    $("#myfooter").append(footer);
});
//获取登录信息,并显示。
function getDataFromBackend() {
    var isLogin = $.get(URL_TASK, {Command:CMD_LOAD_PARA}, function(data){
        var login = "";
        $.each(data, function(key, value){
            if(key == KEY_LOGIN) {
                login = value;
            } else if (key == KEY_WORKER) {
                gLoginInfo.workerId = value;
            } else if (key == KEY_TITLE) {
                gLoginInfo.title = value;
                console.log("title =" + gLoginInfo.title);
            }
        });

        if (login != STR_ON) {
            window.location.assign("./login.html");
        } else {
            console.log("login on");
            //显示登录信息
            var Info ='您好，'+ gLoginInfo.workerId+ ',<a href="javascript:void(0);" onclick="dropoutpage()">注销</a>';
            var pageTitle = $('title').text();
            $("#login_info").append(Info);
            if ((gLoginInfo.title == STR_ADMIN)&&(pageTitle != "价格管理")){
                var priceMenu = '<li><a href="./price.html">价格管理</a></li>'
                //显示price Menu
                $("#my_menu").append(priceMenu);
            }
        }
    });
    return isLogin;
}
//退出登录
function dropoutpage() {
    $.post(URL_TASK, {Command:CMD_UNLOAD}, function(){
        window.location.assign("./login.html");
    });
}
//毫秒转换为yyyy/mm/dd
function timeToDate(value) {
    if(value != 0){
        var t = new Date(value * 1000);
        return (t.getFullYear() + "-" + (t.getMonth() + 1) + "-" + t.getDate());
    }else{
        return 0;
    }

}
