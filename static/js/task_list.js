var login;
var workerId;
var title;
var tasks = new Array();

$(document).ready(function(){
    getDataFromBackend();
    getTaskList(); /** 获取全部任务信息，并以table形式来显示**/
    $("#myModal").on("hidden.bs.modal", function() {
        $("#detail_show").empty();
        $("#detail_show").hide();
        $("#task_form").reset;
        $("#task_form").hide();
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
      displayWorkerId(login,workerId);
  });
}
/** 获取全部任务信息**/
function getTaskList() {
    var today = new Date();
    var stime = Date.UTC(today.getFullYear(), today.getMonth(), today.getDay(), 0, 0, 0, 0);
    console.log("get task list! time=" + stime/1000);

    $.get(URL_TASK, {Command:CMD_LOAD_TASK, StartTime:0}, function(data){
        $.each(data, function(key, value){
            console.log("key=" + key + ", value=" + value);
            if (key == KEY_TASKS) {
                $.each(value, function(index, obj){
                    tasks[index] = obj;
                    descriptionTask(obj);
                });
            }
        });
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
        console.log("item =" + item + ", value=" + task[item]);
        if((item == "SponsorId")||(item == "CreateTime")||(item == "Type")||(item == "State")) {
            task_info = task_info + "<td>" + task[item] + "</td>";
        } else if(item == "TaskId"){
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