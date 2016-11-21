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
    displayFooter();
    displayHeader();
    $.when(getDataFromBackend()).done(function(){
        getTaskList();
        console.log("first login=",gLoginInfo.title);
        displayTaskAction();
    });
    bindMyModalClick();
    searchAction();
    initdate();
});
//绑定模态框关闭事件
function bindMyModalClick(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#detail_win").empty().hide();
        $("#task_form").hide();
        $("#task_form input").val("");
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
        btnAction.CancelTask();
        btnAction.AssignTask();
        btnAction.CheckTask();
        btnAction.CommitTask();
        btnAction.CloseTask();
        btnAction.ArchiveTask();
    }
}
//重新加载任务列表
function reLoadTasksList() {
    $("#task_list").empty();
    $("#task_list input[type=checkbox]").prop("checked",false);
    selectedTasks.length=0;
    getTaskList();
}
//转换task信息
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

//获取全部任务信息
function getTaskList(){
    tasks.length = 0;
    $.get(URL_TASK, {Command:CMD_LOAD_TASK, STime:0, ETime:0, Worker:"", State:"", Farm:"", Cell:"", Patch:""},
        function(data){
            $.each(data, function(key, value){
                console.log("key=" + key + ", value=" + value);
                if (key == KEY_TASKS) {
                    $.each(value, function(index, obj){
                        var item = getTaskListItem(obj);
                        tasks.push(item);
                    });
                }
            });
        }).done(function () {
        descriptionTask(tasks);
    });
}

//解析全部任务信息
function descriptionTask(data) {
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
        task_info = "<tr>"+task_info+"</tr>";
        //增加任务行
        //console.log("task_info=" + task_info);
        $("#task_list").append(task_info);
    });
}
//转换log信息
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

        var logInfo ='<li>'+actiontime+','+operate+','+action+','+comment+'</li>';



    return logInfo;

    //console.log("Print Log: Id=" + id + ", Action=" + action + ", Operate=" + operate + ", ActionTime=" + actiontime + ", comment=" + comment);
}
//查询并显示任务详情,以及log信息
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
                        console.log("logInfo:" + logInfo);
                    });
                }
            });
    }).done(function () {
            console.log("logInfo=" + logInfo);

            logInfo = '<ul>'+ logInfo+ '</ul>'

            console.log("logInfo=" + logInfo);
            var obj=tasks[index-1];
            for(item in obj){
                task_details_info +="<tr><td>"+item+"</td><td>"+obj[item]+"</td><tr>";
            }
            task_details_info +='<tr>'+'<td colspan="2" style="text-align: left"><label>logInfo:</label>'+logInfo+'</td></tr>';
            task_details_info= '<table class="table table-bordered table-hover table-condensed"><tbody>'+
                task_details_info+'<tbody></table>';

            $("#myModalLabel").text("任务详情");
            $("#detail_win").show().append(task_details_info); //显示详情模态框内容
    });
}

//任务的各种action
var btnAction={
    //创建任务
    CreateTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CreateTask" data-toggle="modal" data-target="#myModal">添加</button>');
        $("#CreateTask").click(function () {
            $("#task_form").show().reset;
            $("#myModalLabel").text("新建任务");
            $("#modesavebtn").show().on("click",function () {
                var farmid = $("#task-farm").val();
                var cellid = $("#task-cell").val();
                var patchid = $("#task-patch").val();
                var workerid = $("#task-worker").val();
                var type = $("#task-type").val();
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
            if(isChecked()) {
                var item = new SelectedCmdParaConstructor();
                console.log("delete cmd parameter = " + item);
                var json = JSON.stringify(item);
                console.log("json string = " + json);
                $.get(URL_TASK, {Command: CMD_CANCEL_TASK, CmdPara: json}, function (data) {
                    console.log("");
                    $.each(data, function (key, value) {
                        if (key == KEY_ERRCODE) {
                            errcode = value;
                        }
                    });
                    if (errcode == 1) {
                        alert("删除失败");
                    } else {
                        alert("删除成功");
                        reLoadTasksList();
                    }
                });
            }
        });
    },
//分配任务
    AssignTask:function () {
        getWorkers();//获取员工Id
        $("#taskBtn").append('<button class="btn btn-default" id="AssignTask">分配</button>');
        $("#AssignTask").click(function () {
            if(isChecked()) {
                $("#myModal").modal("show");
                $("#myModalLabel").text("分配任务");
                $("#assign_win").show();
                $("#modesavebtn").show().on("click", function () {
                    var worker = $("#AssignWorker").val();
                    var checker = $("#AssignChecker").val();
                    worker = worker.substring(0, worker.indexOf("|"));
                    checker = checker.substring(0, checker.indexOf("|"));
                    var item = new AssignCmdParaConstructor(worker, checker);
                    console.log("delete cmd parameter = " + item);
                    var json = JSON.stringify(item);
                    console.log("json string = " + json);
                    $.get(URL_TASK, {Command: CMD_ASSIGN_TASK, CmdPara: json}, function (data) {
                        console.log("");
                        $.each(data, function (key, value) {
                            if (key == KEY_ERRCODE) {
                                errcode = value;
                            }
                        });
                        if (errcode == 1) {
                            alert("分配失败");
                        } else {
                            alert("分配成功");
                            reLoadTasksList();
                            $("#myModal").modal("hide")
                        }
                    });
                });
            }
        });
    },
//提交任务
    CommitTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CommitTask">提交</button>');
        $("#CommitTask").click(function () {
            // test, send the delete command
            if(isChecked()) {
                var item = new SelectedCmdParaConstructor();
                console.log("delete cmd parameter = " + item);
                var json = JSON.stringify(item);
                console.log("json string = " + json);
                $.get(URL_TASK, {Command: CMD_SUBMIT_TASK, CmdPara: json}, function (data) {
                    console.log("");
                    $.each(data, function (key, value) {
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
            }
        });
    },
//检查任务
    CheckTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CheckTask">检查</button>');
        $("#CheckTask").click(function () {
            // test, send the delete command
            if(isChecked()) {
                var item = new SelectedCmdParaConstructor();
                console.log("delete cmd parameter = " + item);
                var json = JSON.stringify(item);
                console.log("json string = " + json);
                $.get(URL_TASK, {Command: CMD_CHECK_TASK, CmdPara: json}, function (data) {
                    console.log("");
                    $.each(data, function (key, value) {
                        if (key == KEY_ERRCODE) {
                            errcode = value;
                        }
                    });
                    if (errcode == 1) {
                        alert("检查失败");
                    } else {
                        alert("检查成功");
                        reLoadTasksList();
                    }
                });
            }
        });
    },
//归档任务
    ArchiveTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="ArchiveTask">归档</button>');
        $("#ArchiveTask").click(function () {
            // test, send the delete command
            if(isChecked()) {
                var item = new SelectedCmdParaConstructor();
                console.log("delete cmd parameter = " + item);
                var json = JSON.stringify(item);
                console.log("json string = " + json);
                $.get(URL_TASK, {Command: CMD_ARCHIVE_TASK, CmdPara: json}, function (data) {
                    console.log("");
                    $.each(data, function (key, value) {
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
            }
        });
    },
//关闭任务
    CloseTask:function () {
        $("#taskBtn").append('<button class="btn btn-default" id="CloseTask">关闭</button>');
        $("#CloseTask").click(function () {
            // test, send the delete command
            if(isChecked()) {
                var item = new SelectedCmdParaConstructor();
                console.log("delete cmd parameter = " + item);
                var json = JSON.stringify(item);
                console.log("json string = " + json);
                $.get(URL_TASK, {Command: CMD_CLOSE_TASK, CmdPara: json}, function (data) {
                    console.log("");
                    $.each(data, function (key, value) {
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
            }
        });
    }

}
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
        }).done(function () {
            reLoadTasksList();
        });
    });
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
//检查是否有选中任务
function isChecked() {
    if(selectedTasks.length > 0){
        return true;
    }else{
        alert("请选择任务");
        return false;
    }
    
}
//获取全部用户信息,放到分配任务下拉框里
function getWorkers() {
    var workersId = new Array();
    var checkerId = new Array();

    $.get(URL_WORKER, {Command:CMD_LOAD_WORKER}, function(data){
        $.each(data, function(key, value){
            if (key == KEY_WORKERS) {
                $.each(value, function(index, obj){
                    var worker;
                    var title;
                    var name;
                    for(item in obj){
                        if (obj[KEY_CHECKOUTTIME]==0) {
                            worker = obj[KEY_WORKERID];
                            title = obj[KEY_TITLE];
                            name = "|" + obj[KEY_NAME];
                        }
                    }
                    if(title == STR_MANAGER){
                        checkerId.push(worker+name);
                    }
                    else if(title == STR_WORKER){
                        workersId.push(worker+name);
                    }
                });
            }
        });
    }).done(function () {
        var  AssignWorkerInfo = '<option></option>';
        var AssignCheckerInfo ='<option></option>';
        console.log("checkerId:"+checkerId);
        for(var i=0;i<workersId.length;i++){
            AssignWorkerInfo += '<option>'+workersId[i]+'</option>';
        };
        for(var j=0;j<checkerId.length;j++){
            AssignCheckerInfo += '<option>'+checkerId[j]+'</option>';
        };

        AssignWorkerInfo ='<div>工人：<select class="form-control" id ="AssignWorker">'+
            AssignWorkerInfo+
            '</select></div>';

        AssignCheckerInfo ='<div>检查员：<select class="form-control" id ="AssignChecker">'+
            AssignCheckerInfo+
            '</select></div>';

        $("#assign_win").append(AssignWorkerInfo + AssignCheckerInfo);
    });
}
