var KEY_LOGIN = "Login";
var KEY_WORKER = "Worker";
var KEY_TITLE = "Title";
var KEY_FARM = "Farm";
var KEY_TASKS = "Tasks";

var CMD_LOAD_PARA = "LoadPara";
var CMD_LOAD_TASK = "LoadTask";
var CMD_UNLOAD = "UnLoad";

var SEVER_URL = "/land/worker";

var login;
var workerId;
var title;
var tasks = new Array();

function getDataFromBackend() {
  $.get(SEVER_URL, {Command:CMD_LOAD_PARA}, function(data){
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
    for ( item in task) {
        console.log("item =" + item + ", value=" + task[item]);
    }
}

function getTaskList() {
    var today = new Date();
    var stime = Date.UTC(today.getFullYear(), today.getMonth(), today.getDay(), 0, 0, 0, 0);
    console.log("get task list! time=" + stime/1000);

    $.get(SEVER_URL, {Command:CMD_LOAD_TASK, StartTime:0}, function(data){
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

$(window).unload(function(){ 
    //alert("获取到了页面要关闭的事件了！"); 
    console.log("onfunc unload the window");
    $.post("/land/worker", {Command:CMD_UNLOAD}, function(data){
    });    
});