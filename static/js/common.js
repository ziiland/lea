//// key constant
var KEY_LOGIN = "Login";
var KEY_WORKER = "Worker";
var KEY_WORKERS = "Workers";
var KEY_TITLE = "Title";
var KEY_FARM = "Farm";
var KEY_TASKS = "Tasks";
var KEY_COMMAND = "Command";
var KEY_PASSWORD = "Password";
var KEY_LOGS = "Logs";

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


// command
var CMD_LOAD_PARA = "LoadPara";
var CMD_LOAD_TASK = "LoadTask";
var CMD_LOAD_WORKER = "LoadWorker";
var CMD_UNLOAD = "UnLoad";
var CMD_LOGIN = "Login";
var CMD_ADD_WORKER = "AddWorker"
var CMD_QUERY_TASK = "QueryTask";

// url
var URL_LOGIN = "/land/login";
var URL_WORKER = "/land/worker";
var URL_TASK = "/land/task";

// event
var EVT_HIDE_BTNADD = "HideBtnAdd";
var EVT_PARA_LOADED = "ParaLoaded";
var EVT_TASKS_LOADED = "TaskLoaded";

//
var gTaskStateDes = {Created:"已创建", Assigned:"已分配", Started:"进行中", Finished:"已完成", Checked:"已检查", Closed:"已关闭", Archived:"已归档"};
var gTaskTypes = ["翻地", "播种", "浇水", "施肥", "搭架子", "移栽", "嫁接", "除草", "除虫", "收割", "快递"];

///////////////////////////////////////////////////////////////////////////////
$(document).ready(function(){
    displayFooter();
    displayHeader();
});

/** 显示页头 **/
function  displayHeader() {
    var herder =  '<h1 style="text-align:center">Ziiland生产管理软件 <small>v1.0</small></h1>';

    $("#myherder").append(herder);
}

/** 显示页脚 **/
function  displayFooter() {
    var footer = '<div class="col-md-12 column text-center"> ' +
                    '<h5>版权信息：上海自然技术有限公司</h5> ' +
                '</div>';

    $("#myfooter").append(footer);
}

/** 显示登录的用户 **/
function displayWorkerId(login,workerid) {
    if (login != "on") {
        window.location.assign("./login.html");
    } else {
        var Info ='<label>您好，'+ workerid+'</label>'+
            '<button class="btn btn-sm" onclick="dropoutpage()">注销</button>';

        $("#login_info").append(Info);
    }
}

/** 退出登录 **/
function dropoutpage() {
    $.post(URL_TASK, {Command:CMD_UNLOAD}, function(){
        window.location.assign("./login.html");
    });
}

