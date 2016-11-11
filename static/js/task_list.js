var tasks = new Array();
var selectedTasks = new Array();
///////////////////////////////////////////////////////////////////////////////
// core-object constructor
function TaskConstructor(taskId, sponsorId, farmId, cellId, patchId, workerId, checkerId,
                         state, type, createTime, startTime, endTime, checkTime, score, userComment, comment){
    this.TaskId = taskId;
    this.SponsorId = sponsorId;
    this.FarmId = farmId;
    this.CellId = cellId;
    this.PatchId = patchId;
    this.WorkerId = workerId;
    this.CheckerId = checkerId;
    this.State = state;
    this.Type = type;
    this.CreateTime = createTime;
    this.StartTime = startTime;
    this.EndTime = endTime;
    this.CheckTime = checkTime;
    this.Score = score;
    this.UserComment = userComment;
    this.Comment = comment;
}

//function TaskSearch
function DeleteCmdParaConstructor() {
    // delete the selected tasks
    this.Tasks = selectedTasks;
}

///////////////////////////////////////////////////////////////////////////////
$(document).ready(function(){
    getDataFromBackend();
    displayFooter();
    displayHeader();
    $(document).on(EVT_PARA_LOADED, function(){
        taskOper.getTaskList();
    });

    $(document).on(EVT_TASKS_LOADED, function() {
        taskOper.descriptionTask(tasks);
    });
    bindMyModalClick();
    displayTaskAction();

    $("#form_search").submit(function(event){
        event.preventDefault();
        //DoSendLoginInfo();
        console.log("Press the search button");
    });
});

//绑定各种click
function bindMyModalClick(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#detail_show").empty().hide();
        $("#task_form").hide();
        $("#modesavebtn").hide();
    });
}
//task action
function displayTaskAction() {
    btnAction.addTask();
}

function getTaskListItem(data) {
    var taskId, sponsorId, farmId, cellId, patchId, workerId, checkerId;
    var state, type, createTime, startTime, endTime, checkTime, score, userComment, comment;
    $.each(data, function(key, value){
        if (key == KEY_TASK_TASKID) {
            taskId = value;
        } else if(key == KEY_TASK_SPONSORID) {
            sponsorId = value;
        } else if(key == KEY_TASK_FARMID) {
            farmId = value;
        } else if(key == KEY_TASK_CELLID) {
            cellId = value;
        } else if(key == KEY_TASK_PATCHID) {
            patchId = value;
        } else if(key == KEY_TASK_WORKERID) {
            workerId = value;
        } else if(key == KEY_TASK_CHECKERID) {
            checkerId = value;
        } else if(key == KEY_TASK_STATE) {
            state = value;
        } else if(key == KEY_TASK_TYPE) {
            type = value;
        } else if(key == KEY_TASK_CREATETIME) {
            createTime = value;
        } else if(key == KEY_TASK_STARTTIME) {
            startTime = value;
        } else if(key == KEY_TASK_ENDTIME) {
            endTime = value;
        } else if(key == KEY_TASK_CHECKTIME) {
            checkTime = value;
        } else if(key == KEY_TASK_SCORE) {
            score = value;
        } else if(key == KEY_TASK_UCOMMENT) {
            userComment = value;
        } else if(key == KEY_TASK_COMMENT) {
            comment = value;
        } else {
            // error
        }
    });

    var item = new TaskConstructor(taskId, sponsorId, farmId, cellId, patchId, workerId, checkerId,
        state, type, createTime, startTime, endTime, checkTime, score, userComment, comment);
    return item;
}
var taskOper = {
    //获取全部任务信息
    getTaskList:function () {

        $.get(URL_TASK, {Command:CMD_LOAD_TASK, STime:0, ETime:0, Worker:"", State:"", Farm:"SHA001", Cell:"", Patch:""}, function(data){
                $.each(data, function(key, value){
                    console.log("key=" + key + ", value=" + value);
                    if (key == KEY_TASKS) {
                        $.each(value, function(index, obj){
                            var item = getTaskListItem(obj);
                            tasks.push(item);
                        });
                    }
                });
                $(document).trigger(EVT_TASKS_LOADED);
            });
        },

    //解析全部任务信息
    descriptionTask:function(data) {
        $.each(data, function(index, value){
            var task_info='<td><input type="checkbox" onclick="setCheckedId(this)"></td>';
            for (item in value) {
                if((item == KEY_TASK_TASKID) ||(item == KEY_TASK_FARMID)|| 
                    (item == KEY_TASK_WORKERID)||(item == KEY_TASK_CHECKERID) || (item == KEY_TASK_COMMENT)) {
                    task_info += "<td>" + value[item] + "</td>";
                } else if(item == KEY_TASK_STATE){
                    //console.log("Task state=" + value[item] + "end");
                    task_info += "<td>" + gTaskStateDes[value[item]] + "</td>";
                } else if(item == KEY_TASK_TYPE) {
                    //var id = value[item];
                    //console.log("id=" +id);
                    task_info += "<td>" + gTaskTypes[value[item]] + "</td>";
                }
            }
            task_info =task_info+"<td><button class='btn btn-sm btn-info' onclick='TaskDetailsAction(this)' data-toggle='modal' data-target='#myModal'>详情</button></td>";
            task_info = "<tr>"+task_info+"</tr>"
            /** 增加任务行*/
            $("#tasklist").append(task_info);
        });
    }
}

function PrintLog(data) {
    var action="";
    var operate="";
    var actiontime=";"
    var comment="";

    $.each(data, function(key, value){
        //console.log("Print log: key=" + key + ", value=" + value);
        if (key == KEY_LOG_ACTION) {
            action = value;
        } else if (key == KEY_LOG_OPERATORID) {
            operate = value;
        } else if (key == KEY_LOG_ACTIONTIME) {
            actiontime = value;
        } else if (key == KEY_TASK_COMMENT) {
            comment = value;
        }
    });
    var logInfo ='<tr><td>'+action+'</td><td>'+operate+'</td><td>'+actiontime+'</td><td>'+comment+'</td></tr>';

    return logInfo;

    console.log("Print Log: Id=" + id + ", Action=" + action + ", Operate=" + operate + ", ActionTime=" + actiontime + ", comment=" + comment);
}

//查询并显示任务详情
function TaskDetailsAction(o){
    var task_details_info = "";
    var index = o.parentNode.parentNode.rowIndex;
    var task_id = $(o).parent().next().text();
    var logInfo ="";
    console.log("index=" + index);
    // get task log
    $.get(URL_TASK, {Command:CMD_QUERY_TASK, TaskId:task_id}, function(data){
        $.each(data, function(key, value){
            if (key == KEY_LOGS)  {
                $.each(value, function(index, obj){
                    // logs
                    logInfo += PrintLog(obj);
                });
            }
        });
    });
    logInfo = '<table class="table table-bordered table-hover table-condensed">' +
                '<thead><th>action</th>' +
                '<th>operate</th>'+
                '<th>actiontime</th>'+
                '<th>comment</th>'+
                '</thead><tbody>'+ logInfo+'<tbody></table>';


    var obj=tasks[index-1];
        for(item in obj){
            task_details_info +="<tr><td>"+item+"</td><td>"+obj[item]+"</td><tr>";
        }
       task_details_info +='<tr>'+logInfo+'</tr>';
        task_details_info= '<table class="table table-bordered table-hover table-condensed"><tbody>'+
                             task_details_info+'<tbody></table>';

        $("#myModalLabel").text("任务详情");
        $("#detail_show").show().append(task_details_info)//显示详情模态框内容
}

//任务的各种action
var btnAction={
    //显示创建任务表单
    addTask:function () {
        $("#task-action").append('<button class="btn btn-default" id="addTask" data-toggle="modal" data-target="#myModal">添加</button>');
        $("#addTask").click(function () {
            // test, send the delete command
            var item = new DeleteCmdParaConstructor();
            console.log("delete cmd parameter = " + item);
            var json = JSON.stringify(item);
            console.log("json string = " + json);
             $.get(URL_TASK, {Command:CMD_CANCEL_TASK, CmdPara:json}, function(data){
                console.log("");
             });

            $("#task_form").show().reset;
            $("#myModalLabel").text("新建任务");
            $("#modesavebtn").show();
        });
    },

    dellTask:function () {
        $("#task-action").appendChild('<button class="btn btn-default" id="dellTask">删除</button>');
        $("#dellTask").click(function () {

        });
    },

    allocationTask:function () {
        $("#task-action").appendChild('<button class="btn btn-default" id="allocationTask">分配</button>');
        $("#allocationTask").click(function () {

        });
    },

    commitTask:function () {
        $("#task-action").appendChild('<button class="btn btn-default" id="commitTask">提交</button>');
        $("#commitTask").click(function () {

        });
    },

    checkTask:function () {
        $("#task-action").appendChild('<button class="btn btn-default" id="checkTask">检查</button>');
        $("#checkTask").click(function () {

        });
    },
    //保存，创建任务begin
    modeSaveBtn:function () {
        var farmid=$("#task-farm").val();
        var cellid=$("#task-cell").val();
        var patchid=$("#task-patch").val();
        var workerid=$("#task-worker").val();
        var type=$("#task-type").val();
        var usercomment=$("#user-comment").val();
        var comment=$("#task-comment").val();

        if (farmid != "" && cellid!= "" && patchid!="" && workerid!="" && type!="") {
            $.post(URL_TASK, { Command: CMD_ADD_TASK, FarmId: farmid, CellId: cellid, PatchId: patchid, WorkerId: workerid,
                Type: type, UserComment: usercomment, Comment: comment}, function (data) {
                $.each(data, function(key,value){
                    if (key == "Errcode") {
                        errcode = value;
                    }
                });
                if (errcode == 1) {
                    alert("创建不成功");
                } else {
                    $("#myModal").modal("hide");
                    alert("新增成功");
                }
            });
        } else {
            alert("请输入正确的信息");
        }
    }
    //保存，创建任务end
}

//选中或者取消所有
function choseAllBox(o){
    var taskId = "";
    console.log("length=" + $("#tasks_table").find("tr").length );

    if($(o).prop("checked")==true) {
        $("input[type=checkbox]").prop("checked",true);
        selectedTasks.length=0;//先清空，再添加全部taskid
        for(var i=0;i<$("#tasks_table").find("tr").length;i++) {
            taskId = $("#tasks_table").find("tr").eq(i).find("td:eq(1)").text();
           // console.log("taskId=" + taskId );
            selectedTasks.push(taskId);
            console.log("selectedTasks=" + selectedTasks );
        }
    }else{
        $("input[type=checkbox]").prop("checked",false);
        selectedTasks.length=0;
        console.log("selectedTasks=" + selectedTasks );
    }
}

//checkebox event,如果选中，把taskid放到checkedtask数组里，如果取消选中则删除。
function setCheckedId(data) {
    var taskId = $(data).parent().next().text();
    console.log("taskId=" + taskId );
    if ($(data).prop("checked")) {
        selectedTasks.push(taskId);
    } else {
        var index = data.parentNode.parentNode.rowIndex;
        selectedTasks.splice(index, 1);
        $("#checkAll").removeAttr("checked");
    }
    console.log("selectedTasks=" + selectedTasks );
}