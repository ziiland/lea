var login;
var workerId;
var title;
var tasks = new Array();

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

///////////////////////////////////////////////////////////////////////////////
$(document).ready(function(){
    getDataFromBackend();
     /** 获取全部任务信息，并以table形式来显示**/
    //console.log("tasks = " + tasks);
    $("#myModal").on("hidden.bs.modal", function() {
        $("#detail_show").empty();
        $("#detail_show").hide();
        $("#task_form").reset;
        $("#task_form").hide();
    });

    $(document).on(EVT_PARA_LOADED, function(){
        // already loaded para, now load task list
        console.log("para loaded!");
        getTaskList();
    });

    $(document).on(EVT_TASKS_LOADED, function() {
        // display the task list
        console.log("task data loaded!");
        $.each(tasks, function(index, obj){
            //console.log("tasks: index=" + index + ", obj=" + obj);
            descriptionTask(obj);
        });
    });
});

function getDataFromBackend() {
  $.get(URL_TASK, {Command:CMD_LOAD_PARA}, function(data){
      $.each(data, function(key, value){
          if(key == KEY_LOGIN) {
              login = value;
          } else if (key == KEY_WORKER) {
              workerId = value;
          } else if (key == KEY_TITLE) {
              title = value;
          }
      });

      $(document).trigger(EVT_PARA_LOADED);
      displayWorkerId(login,workerId);
  });
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

/** 获取全部任务信息**/
function getTaskList() {
    var today = new Date();
    var stime = Date.UTC(today.getFullYear(), today.getMonth(), today.getDay(), 0, 0, 0, 0);
    console.log("get task list! time=" + stime/1000);

    $.get(URL_TASK, {Command:CMD_LOAD_TASK, STime:0, ETime:0, Worker:"", State:"Assigned", Farm:"SHA001", Cell:"", Patch:""}, function(data){
        $.each(data, function(key, value){
            console.log("key=" + key + ", value=" + value);
            if (key == KEY_TASKS) {
                $.each(value, function(index, obj){
                    //tasks[index] = obj;
                    //descriptionTask(obj);
                    var item = getTaskListItem(obj);
                    tasks.push(item);
                });
            }
        });

    $(document).trigger(EVT_TASKS_LOADED);    
    });
}

/**解析全部任务对象数组
 *创建任务table
 *绑定每个任务详情按键event
 **/
function descriptionTask(task) {
    var task_info = "";
    var task_id = "";

    for (item in task) {
        //console.log("item =" + item + ", value=" + task[item]);
        if((item == KEY_TASK_SPONSORID) || (item == KEY_TASK_CREATETIME) || (item == KEY_TASK_STATE)) {
            task_info = task_info + "<td>" + task[item] + "</td>";
        } else if(item == KEY_TASK_TYPE) {
            index = task[item];
            console.log("task type = " + index);
            //task_info = task_info + "<td>" + task[item] + "</td>";
            task_info = task_info + "<td>" + gTaskTypes[index] + "</td>";
        } else if(item == KEY_TASK_TASKID){
            task_id = task[item];
        }
    }
    console.log("task_id==="+task_id);
    task_info = "<tr>"+task_info+"<td><button class='btn btn-default' data-toggle='modal' data-target='#myModal' id="+task_id+">详情</button></td></tr>"
    /** 增加任务行*/
    $("#tasklist").append(task_info);
    /** 绑定任务详情click event**/
    $("#"+task_id).click(function () {
        displayDetailsTask(task_id);
    });

}
/**查询并显示任务详情**/
function displayDetailsTask(task_id){
    var task_details_info = "";
    $.get(URL_TASK, {Command:CMD_QUERY_TASK, TaskId:task_id}, function(data){
        $.each(data, function(key, value){
            if (key == KEY_TASKS)  {
                $.each(value, function(index, obj){
                    console.log("obj=" + obj );
                    for(item in obj){
                        task_details_info = task_details_info +"<dt style='color: red'>"+item+":</dt>"+ "<dd>"+obj[item]+"</dd><hr>";
                    }
                    $("#detail_show").show();
                    $("#myModalLabel").text("任务详情");
                    $("#detail_show").append("<dl>"+task_details_info+"</dl>")/** 显示详情模态框内容**/
                });
            }
        });
    });
}
/**显示创建任务表单**/
function displayAddTaskUi(){
    $("#task_form").show();
    $("#myModalLabel").text("新建任务");
}