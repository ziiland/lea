var errcode = 1;
var sender;
var receiver;
var address;
var mobile;

var workerId = "";
var password = "";
$(document).ready(function(){
    //
    // getCookies();
    // AutoCursorMove();
});

function getPacketContent() {
    taskId = $("#TaskId").val();
    console.log("taskId = " + taskId);

    $.get(URL_PACKET, {Command: CMD_GET_PACKET, TaskId:taskId, ExpressNo:"", PacketId:1}, function (data) {
        $.each(data, function (key, value) {
            if (key == "Errcode") {
                errcode = value;
            } else if (key == "Sender") {
                sender = value;
            } else if (key == "Receiver") {
                receiver = value;
            } else if (key == "Address") {
                address = value;
            } else if (key == "RMobile") {
                mobile = value;
            }
        });
        console.log("errcode=" + errcode + 
            ", Sender=" + sender + 
            ", Receiver=" + receiver + 
            ", Address=" + address + 
            ", RMobile=" + mobile);
        // if (errcode == 0 && page != "") {
        //     // goto the page
        //     window.location.assign("./" + page);
        // } else {
        //     $("#WorkerId").val("");
        //     $("#Password").val("");
        //     workerId = "";
        //     password = "";
        //     alert("请输入正确的用户名和密码");
        // }
    });
}

// function  AutoCursorMove(){
//     $(document).keydown(function (event) {
//         console.log("event.keyCode="+event.keyCode);
//         var focus_id=document.activeElement.id;
//         if(event.keyCode==13){
//             if(focus_id=="WorkerId"){
//                 $("#Password").focus()
//             }else if(focus_id=="Password"){
//                 DoSendLoginInfo();
//             }
//         }
//     });
// }
// //获取用户名密码
// function getCookies() {
//     if (localStorage.getItem("rmbUser") == "true") {
//         $("#remall").prop("checked", true);
//         $("#WorkerId").val(localStorage.getItem("username"));
//         $("#Password").val(localStorage.getItem("password"));
//         $("#sign_button").focus();
//     }
//     else {
//         $("#WorkerId").focus();
//     }
// }

// //记住用户名密码
// function save() {
//     if ($("#remall").prop("checked")) {
//         var username = $("#WorkerId").val();
//         var password = $("#Password").val();
//         localStorage.setItem("rmbUser", "true"); //存储一个带7天期限的cookie
//         localStorage.setItem("username", username);
//         localStorage.setItem("password", password);
//     }else{
//         localStorage.setItem("rmbUser", "false");
//         localStorage.setItem("username", "");
//         localStorage.setItem("password", "");
//     }
// }

// function DoSendLoginInfo() {
//     workerId=$("#WorkerId").val();
//     password=$("#Password").val();
//     console.log("workerId=" + workerId + ", password=" + password);
//     if((workerId =="")&&(password != "")){
//         $("#WorkerId").focus()
//         alert("请输入用户名");
//     }
//     else if((workerId !="")&&(password == "")){
//         $("#Password").focus()
//         alert("请输入密码");
//     }else if((workerId =="")&&(password == "")){
//         $("#WorkerId").focus()
//         alert("请输入用户名和密码");
//     }
//     if((workerId!="")&&(password!="")){
//         save();
//         $.get(URL_LOGIN, {Command: CMD_LOGIN, Worker: workerId, Password: password}, function (data) {
//             $.each(data, function (key, value) {
//                 if (key == "Errcode") {
//                     errcode = value;
//                 } else if (key == "Page") {
//                     page = value;
//                 }
//             });
//             console.log("errcode=" + errcode + ", page=" + page);
//             if (errcode == 0 && page != "") {
//                 // goto the page
//                 window.location.assign("./" + page);
//             } else {
//                 $("#WorkerId").val("");
//                 $("#Password").val("");
//                 workerId = "";
//                 password = "";
//                 alert("请输入正确的用户名和密码");
//             }
//         });
//     }
// }

