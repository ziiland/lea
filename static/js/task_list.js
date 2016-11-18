var tasks = new Array();
var selectedTasks = new Array();
var errcode = 1;
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

function AssignCmdParaConstructor(worker, checker) {
    this.Tasks = selectedTasks;
    this.Worker = worker;        // use workerId
    this.Checker = checker;      // use checkerId
}

//selected Constructor
function SelectedCmdParaConstructor() {
    // delete the selected tasks
    this.Tasks = selectedTasks;
}

function SearchCmdParaConstructor() {
    // 
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
        bindMyModalClick();
        displayTaskAction();
        searchAction();
        initdate();
    });
});
//搜索按钮事件
function searchAction() {
    $("#searchBtn").click(function () {
        console.log("Press the search button");
        var stime = $("#starttime").val();
        var etime = $("#endtime").val();
        var worker = $("#searchworkerid").val();
        var state = $("#searchtype").val();
        var farm = $("#searchfarmid").val();

        console.log("type: stime=" + typeof(stime) + ", etime=" + typeof(etime) + ", worker=" + typeof(worker)
            + ", state=" + typeof(state) + ", farm=" + typeof(farm));

        tasks.length = 0;
        $("#tasklist").empty();
        $.get(URL_TASK, {Command:CMD_LOAD_TASK, STime:stime, ETime:etime, Worker:worker, State:state, Farm:farm, Cell:"", Patch:""}, function(data){
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
    });
}
//绑定模态框关闭事件
function bindMyModalClick(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#detailwindown").empty().hide();
        $("#task_form").hide();
        $("#modesavebtn").hide().off("click");
    });
}
//时间控件
function initdate() {
    $('#starttime').fdatepicker();
    $('#endtime').fdatepicker();
}
//display tasksaction button
function displayTaskAction() {
    console.log("displayTaskAction title="+gLoginInfo.title);
    if(gLoginInfo.title == STR_WORKER) {
        btnAction.CommitTask();
    } else if (gLoginInfo.title == STR_ADMIN || gLoginInfo.title == STR_MANAGER){
        btnAction.CreateTask();
        btnAction.ArchiveTask();
        btnAction.AssignTask();
        btnAction.CheckTask();
        btnAction.CloseTask();
        btnAction.CancelTask();
        btnAction.CommitTask();
    }
}

//重新加载任务列表
function reLoadTasksList() {
    taskOper.getTaskList();
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
        tasks.length = 0;
        $.get(URL_TASK, {Command:CMD_LOAD_TASK, STime:0, ETime:0, Worker:"", State:"", Farm:"", Cell:"", Patch:""}, function(data){
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
            task_info = task_info+"<td><button class='btn btn-sm btn-info' onclick='TaskDetailsAction(this)' data-toggle='modal' data-target='#myModal'>详情</button></td>";
            task_info = "<tr>"+task_info+"</tr>"
            /** 增加任务行*/
            $("#tasklist").append(task_info);
        });
    }
}

function PrintLog(data) {
    var action = "";
    var operate = "";
    var actiontime = "";
    var comment = "";

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

    //console.log("Print Log: Id=" + id + ", Action=" + action + ", Operate=" + operate + ", ActionTime=" + actiontime + ", comment=" + comment);
}

//查询并显示任务详情
function TaskDetailsAction(o){
    var task_details_info = "";
    var index = o.parentNode.parentNode.rowIndex;
    var task_id = $(o).parent().parent().children("td").eq(1).text();
    var logInfo ="";
    console.log("task_id=" + task_id);
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
    }).done(function () {
            console.log("logInfo=" + logInfo);

            logInfo = '<table class="table table-bordered table-hover table-condensed">' +
                '<thead><th>action</th><th>operate</th><th>actiontime</th><th>comment</th></thead>' +
                '<tbody>'+ logInfo +'</tbody>' +
                '</table>';

            console.log("logInfo=" + logInfo);
            var obj=tasks[index-1];
            for(item in obj){
                task_details_info +="<tr><td>"+item+"</td><td>"+obj[item]+"</td><tr>";
            }
            task_details_info +='<tr>'+logInfo+'</tr>';
            task_details_info= '<table class="table table-bordered table-hover table-condensed"><tbody>'+
                task_details_info+'<tbody></table>';

            $("#myModalLabel").text("任务详情");
            $("#detailwindown").show().append(task_details_info); //显示详情模态框内容
    });
}

//任务的各种action
var btnAction={
    //创建任务
    CreateTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CreateTask" data-toggle="modal" data-target="#myModal">添加</button>');
        $("#CreateTask").click(function () {
            // test, send the assign command
            // var item = new AssignCmdParaConstructor("ZLD00004", "ZLD00003");
            // console.log("Assign cmd parameter = " + item);
            // var json = JSON.stringify(item);
            // console.log("json string = " + json);
            // $.get(URL_TASK, {Command:CMD_ASSIGN_TASK, CmdPara:json}, function(data){
            //     console.log("");
            // });

            $("#task_form").show().reset;
            $("#myModalLabel").text("新建任务");
            $("#modesavebtn").show().on("click",function () {
                var farmid = $("#task-farm").val();
                var cellid = $("#task-cell").val();
                var patchid = $("#task-patch").val();
                var workerid = $("#task-worker").val();
                var type = $("#task-type").val();
                var usercomment = $("#user-comment").val();
                var comment = $("#task-comment").val();

                //console.log("Press SaveBtn: type=", typeof(type));
                if (farmid != "" && cellid != "" && patchid != "" && type != "") {
                    $.post(URL_TASK, { Command: CMD_ADD_TASK, Farm: farmid.toUpperCase(),
                        Cell: cellid.toUpperCase(), Patch: patchid.toUpperCase(), 
                        Worker: workerid.toUpperCase(),
                        Type: type, Comment: comment}, function (data) {
                        $.each(data, function(key, value){
                            if (key == KEY_ERRCODE) {
                                errcode = value;
                            }
                        });
                        if (errcode == 1) {
                            alert("创建不成功");
                        } else {
                            $("#myModal").modal("hide");
                            alert("新增成功");
                            reLoadTasksList();
                        }
                    });
                } else {
                    alert("请输入正确的信息");
                }
            });
        })
    },
//删除任务
    CancelTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CancelTask">删除</button>');
        $("#CancelTask").click(function () {
            // test, send the delete command
             var item = new SelectedCmdParaConstructor();
             console.log("delete cmd parameter = " + item);
             var json = JSON.stringify(item);
             console.log("json string = " + json);
             $.get(URL_TASK, {Command:CMD_CANCEL_TASK, CmdPara:json}, function(data){
                    console.log("");
                 $.each(data, function(key, value){
                     if (key == KEY_ERRCODE) {
                         errcode = value;
                     }
                 });
                 if (errcode == 1) {
                     alert("删除不成功");
                 } else {
                     alert("删除成功");
                     reLoadTasksList();
                 }
            });
        });
    },

    AssignTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="AssignTask">分配</button>');
        $("#AssignTask").click(function () {
            $("#task_form").show().reset;
            $("#myModalLabel").text("新建任务");
            $("#modesavebtn").show().on("click",function () {
            });
        });
    },

    CommitTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CommitTask">提交</button>');
        $("#CommitTask").click(function () {

        });
    },

    CheckTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CheckTask">检查</button>');
        $("#CheckTask").click(function () {

        });
    },
    ArchiveTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="ArchiveTask">归档</button>');
        $("#ArchiveTask").click(function () {
            // test, send the delete command
            var item = new SelectedCmdParaConstructor();
            console.log("delete cmd parameter = " + item);
            var json = JSON.stringify(item);
            console.log("json string = " + json);
            $.get(URL_TASK, {Command:CMD_CANCEL_TASK, CmdPara:json}, function(data){
                console.log("");
                $.each(data, function(key, value){
                    if (key == KEY_ERRCODE) {
                        errcode = value;
                    }
                });
                if (errcode == 1) {
                    alert("归档失败");
                } else {
                    alert("归档成功");
                    reLoadTasksList();
                }
            });
        });
    },
    CloseTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CloseTask">关闭</button>');
        $("#CloseTask").click(function () {
            // test, send the delete command
            var item = new SelectedCmdParaConstructor();
            console.log("delete cmd parameter = " + item);
            var json = JSON.stringify(item);
            console.log("json string = " + json);
            $.get(URL_TASK, {Command:CMD_CANCEL_TASK, CmdPara:json}, function(data){
                console.log("");
                $.each(data, function(key, value){
                    if (key == KEY_ERRCODE) {
                        errcode = value;
                    }
                });
                if (errcode == 1) {
                    alert("关闭失败");
                } else {
                    alert("关闭成功");
                    reLoadTasksList();
                }
            });
        });
    }

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