var errcode = 1;
var page = "";

var workerId = "";
var password = "";
$(document).ready(function(){
    displayFooter();
    displayHeader();
    getCookies();
    AutoCursorMove();
});



function  AutoCursorMove(){
    $(document).keydown(function (event) {
        console.log("event.keyCode="+event.keyCode);
        var focus_id=document.activeElement.id;
        if(event.keyCode==13){
            if(focus_id=="WorkerId"){
                $("#Password").focus()
            }else if(focus_id=="Password"){
                DoSendLoginInfo();
            }
        }
    });
}
//��ȡ�û�������
function getCookies() {
    if (localStorage.getItem("rmbUser") == "true") {
        $("#remall").prop("checked", true);
        $("#WorkerId").val(localStorage.getItem("username"));
        $("#Password").val(localStorage.getItem("password"));
        $("#sign_button").focus();
    }
    else {
        $("#WorkerId").focus();
    }
}

//��ס�û�������
function save() {
    if ($("#remall").prop("checked")) {
        var username = $("#WorkerId").val();
        var password = $("#Password").val();
        localStorage.setItem("rmbUser", "true"); //�洢һ����7�����޵�cookie
        localStorage.setItem("username", username);
        localStorage.setItem("password", password);
    }else{
        localStorage.setItem("rmbUser", "false");
        localStorage.setItem("username", "");
        localStorage.setItem("password", "");
    }
}

function DoSendLoginInfo() {
    workerId=$("#WorkerId").val();
    password=$("#Password").val();
    console.log("workerId=" + workerId + ", password=" + password);
    if((workerId =="")&&(password != "")){
        $("#WorkerId").focus()
        alert("�������û���");
    }
    else if((workerId !="")&&(password == "")){
        $("#Password").focus()
        alert("����������");
    }else if((workerId =="")&&(password == "")){
        $("#WorkerId").focus()
        alert("�������û���������");
    }
    if((workerId!="")&&(password!="")){
        $.get(URL_LOGIN, {Command: CMD_LOGIN, Worker: workerId, Password: password}, function (data) {
            $.each(data, function (key, value) {
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
                alert("��������ȷ���û���������");
            }
        });
    }
}

