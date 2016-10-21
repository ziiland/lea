var errcode = 1;
var page = "";

var workerId = "";
var password = "";

function DoSendLoginInfo() {
    //console.log("workerId=" + workerId + ", password=" + password);
    $.get(URL_LOGIN, {Command:CMD_LOGIN, Worker:workerId, Password:password}, function(data){
        $.each(data, function(key,value){
            if (key == "Errcode") {
                errcode = value;
            } else if (key == "Page") {
                page = value;
            }
        });
        console.log("errcode=" + errcode + ", page=" + page);
        if (errcode == 0 && page != "") {
            // goto the page
            window.location.assign("./" + page);
        } else {
            $("#WorkerId").val("");
            $("#Password").val("");
            workerId = "";
            password = "";
        }        
    });  
}

function  AutoCursorMove(){
    $("#WorkerId").change(function(){
    	workerId = $("#WorkerId").val();
        console.log("workerId=" + workerId + ", password=" + password);
        if (workerId != "" && password != "") {
            //DoSendLoginInfo();
        } else {
            $("#WorkerId").blur();
            $("#Password").focus();
        }
    });

    $("#Password").change(function(){
        password = $("#Password").val();
        console.log("workerId=" + workerId + ", password=" + password);
        if (workerId != "" && password != "") {
            //DoSendLoginInfo();
        } else {           
            $("#Password").blur();
            $("#WorkerId").focus();
        }
    });	
}

$(document).ready(function(){
    console.log("ready");
    $("#workerId").focus();    
    AutoCursorMove();

    $("#form_login").submit(function(event){
        event.preventDefault();
        DoSendLoginInfo();
    });
});