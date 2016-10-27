var login;
var workerId;
var title;
var tasks = new Array();

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

      console.log("login=" + login + ", workerId=" + workerId + ", title=" + title);
      if (login != "on") {
          window.location.assign("./login.html");
      }
  });
}

function descriptionTask(task) {

    var task_info ="";
    var task_id="";
    var task_detail_info="";
    for ( item in task) {
        console.log("item =" + item + ", value=" + task[item]);
        if((item == "SponsorId")||(item == "CreateTime")||(item == "Type")||(item == "State")) {
            task_info = task_info + "<td>" + task[item] + "</td>";
        }
        else if(item == "TaskId"){
            task_id = task[item];
        }
        task_detail_info = task_detail_info +"<dt style='color: red'>"+item+":</dt>"+ "<dd>"+task[item]+"</dd><hr>";
    }
    console.log("task_id==="+task_id);
    $("#tasklist").append("<tr>"+task_info+"<td><button class='btn btn-default' id="+task_id+">详情</button></td></tr>");
    $("#"+task_id).click(function () {
       $("#detailModal").modal('show');
       $("#detail_show").append("<dl>"+task_detail_info+"</dl>")
    });
}
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

$(document).ready(function(){
    console.log("ready");
    getDataFromBackend();
    //alert("Load task list!");
    getTaskList();

    // $(window).on("unload", function() {
    //     console.log("onfunc unload the window");
    //     //alert("unload the window");
    //     $.post("/land/worker", {Command:"unload"}, function(data){
    //     });
    // });     
});

// $(window).unload(function(){ 
//     //alert("获取到了页面要关闭的事件了！"); 
//     console.log("onfunc unload the window");
//     $.post(URL_TASK, {Command:CMD_UNLOAD}, function(data){
//     });    
// });