var KEY_LOGIN = "Login";
var KEY_WORKER = "Worker";
var KEY_WORKERS = "Workers";
var KEY_TITLE = "Title";
var KEY_FARM = "Farm";
var KEY_TASKS = "Tasks";
var KEY_COMMAND = "Command";
var KEY_PASSWORD = "Password";

var CMD_LOAD_PARA = "LoadPara";
var CMD_LOAD_TASK = "LoadTask";
var CMD_LOAD_WORKER = "LoadWorker";
var CMD_UNLOAD = "UnLoad";
var CMD_LOGIN = "Login";
var CMD_ADD_WORKER = "AddWorker"
var CMD_QUERY_TASK = "QueryTask";

var URL_LOGIN = "/land/login";
var URL_WORKER = "/land/worker";
var URL_TASK = "/land/task";

var EVT_HIDE_BTNADD = "HideBtnAdd";

$(document).ready(function(){
    displayFooter();
    displayHeader();
});

/** ��ʾҳͷ**/
function  displayHeader() {
    var herder =  '<h1 style="text-align:center">Ziiland����������� <small>v1.0</small></h1>';

    $("#myherder").append(herder);
}

/** ��ʾҳ��**/
function  displayFooter() {
    var footer = '<div class="col-md-12 column text-center"> ' +
                    '<h5>��Ȩ��Ϣ���Ϻ���Ȼ�������޹�˾</h5> ' +
                '</div>';

    $("#myfooter").append(footer);
}
/**x��ʾ��¼���û�**/
function displayWorkerId(login,workerid) {
    if (login != "on") {
        window.location.assign("./login.html");
    } else {
        var Info ='<label>���ã�'+ workerid+'</label>'+
            '<button class="btn btn-sm" onclick="dropoutpage()">ע��</button>';

        $("#login_info").append(Info);
    }
}

/** �˳���¼**/
function dropoutpage() {
    $.post(URL_TASK, {Command:CMD_UNLOAD}, function(){
        window.location.assign("./login.html");
    });
}

