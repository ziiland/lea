var login;
var workerId;

function getDataFromBackend() {
  $.get("/land/worker", {Command:"LoadPara"}, function(data){
      $.each(data, function(key, value){
          if(key == "Login") {
              login = value;
          } else if (key == "WorkerId") {
              workerId = value;
          }
      });

      console.log("login=" + login + ", workerId=" + workerId);
      if (login != "on") {
          window.location.assign("./login.html");
      }
  });
}

$(document).ready(function(){
    console.log("ready");
    getDataFromBackend();

    // $(window).on("unload", function() {
    //     console.log("onfunc unload the window");
    //     //alert("unload the window");
    //     $.post("/land/worker", {Command:"unload"}, function(data){
    //     });
    // });     
});

$(window).unload(function(){ 
    //alert("��ȡ����ҳ��Ҫ�رյ��¼��ˣ�"); 
    console.log("onfunc unload the window");
    $.post("/land/worker", {Command:"unload"}, function(data){
    });    
});