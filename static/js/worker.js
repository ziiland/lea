var workers = new Array();
var gWrokerKey = {"WorkerId":"工号", "Password":"密码", "Name":"姓名"
                , "Sex":"性别", "IdentifyNo":"身份证号", "Title":"角色"
                , "CheckInTime":"入职时间", "CheckOutTime":"离职时间", "Comment":"备注"};
///////////////////////////////////////////////////////////////////////////////
$(document).ready(function(){
    getDataFromBackend();
    displayFooter();//显示页头
    displayHeader();//显示页尾

    $(document).on(EVT_PARA_LOADED, function(){
        getWorkersInfo();//获取全部用户信息
        setModalEvent();//绑定模态框事件
        displayAddButton();
    });
});

//显示添加用户按钮
function displayAddButton() {
    console.log("gLoginInfo.title=displayAddButton_" + gLoginInfo.title)
    if (gLoginInfo.title == STR_ADMIN) {
        $("#add").show(); //显示添加用户按钮
    }
}
//添加用户按钮event
function addClickAction() {
    console.log("Press Add Button!");
    $("#WorkerId").val("");
    $("#Password").val("");
    $("#Name").val("");
    $("#Sex").val("");
    $("#IdentifyNo").val("");
    $("#Title").val("");

    $("#WorkerId").attr("readonly",false);
    $("#registered_from").show();
    $("#myModalLabel").text("新增员工");
    $("#modesavebtn").show();
}

//绑定模态框事件
function setModalEvent(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#worker_detail").empty().hide();
        $("#registered_from").hide().reset;
        $("#modesavebtn").hide();
        $("#modeupdatabtn").hide();
        $("#person_password").hide();//显示详情模态框内容
        $("#modechangebtn").hide();
    });
}
//获取全部用户信息
function getWorkersInfo() {
    workers.length = 0;
    $.get(URL_WORKER, {Command:CMD_LOAD_WORKER}, function(data){
        $.each(data, function(key, value){
            if (key == KEY_WORKERS) {
                $.each(value, function(index, obj){
                    workers[index] = obj;
                    descriptionWorkers(obj);
                });
            }
        });
    });
}

//显示用户信息table
function descriptionWorkers(workers) {
    var worker_info = "";
    var btn_state = 'disabled="disabled"';

    for ( item in workers) {
        //console.log("item =" + item + ", value=" + workers[item]);
        if (item == KEY_TITLE) {
            worker_info = worker_info + "<td>" + gRoleDes[workers[item]] + "</td>";
        } else if(item == KEY_CHECKOUTTIME){
            //console.log("btn_state =" + btn_state);
            if(workers[item] == 0){
                btn_state = "";
            }
            //worker_info = worker_info + "<td>" + workers[item] + "</td>";
        } else if(item == KEY_CHECKINTIME) {
            var checkindate = new Date(workers[item] * 1000);
            // console.log("year = " + checkindate.getFullYear() + ", month = " + (checkindate.getMonth() 
            //     + 1) + ", day = " + checkindate.getDate());
            worker_info += "<td>" + checkindate.getFullYear() + "-" + (checkindate.getMonth() + 1) + "-" 
            + checkindate.getDate() + "</td>";
        } else if ((item != KEY_PASSWORD) && (item != KEY_ID) && (item != "IdentifyNo") && (item != "Comment")) {
            worker_info += "<td>" + workers[item] + "</td>";
        }
    }
    console.log("gLoginInfo.title = "+ gLoginInfo.title)
    console.log("btn_state = " + btn_state);

    var detail_btn='<button class="btn btn-sm btn-info" onclick="workerDetailsAction(this)" data-toggle="modal" data-target="#myModal">详情</button>';
    var del_btn='<button class="btn btn-sm btn-danger"'+btn_state+'onclick="delWorkerAction(this)" >删除</button>';
    var change_btn='<button class="btn btn-sm btn-warning"'+btn_state+' onclick="changeWorkerAction(this)" data-toggle="modal" data-target="#myModal">修改</button>';
    var reset_btn='<button class="btn btn-sm btn-warning" onclick="resetPassword(this)" data-toggle="modal" data-target="#myModal">恢复密码</button>';
    var changePassword_btn='<button class="btn btn-sm btn-info"onclick="changePersonPasswordUi(this)" data-toggle="modal" data-target="#myModal">修改密码</button>';
    if(gLoginInfo.title == STR_ADMIN) {
       if(workers[KEY_WORKERID] == STR_ADMIN) {
           worker_info = worker_info + "<td>" + detail_btn + change_btn + "</td>"
       }else{
           worker_info = worker_info + "<td>" + detail_btn + del_btn+ change_btn + "</td>"
       }
    } else if(gLoginInfo.title == STR_MANAGER){
        if (workers[KEY_WORKERID] == gLoginInfo.workerId.toUpperCase()) {
            worker_info += "<td>" + detail_btn + changePassword_btn + "</td>";
        } else {
            worker_info = worker_info + "<td>" + detail_btn + reset_btn + "</td>"
        }        
    } else {
        worker_info = worker_info +"<td>" + changePassword_btn +"</td>";
    }
    $("#userlist").append("<tr>"+worker_info+"</tr>");
}

//显示用户详情
function workerDetailsAction(o) {
    var worker_details_info = "";
    var index = o.parentNode.parentNode.rowIndex;
    console.log("index=" + index);

    var obj = workers[index-1];
    for(item in obj){
        if (item == KEY_TITLE) {
            worker_details_info += "<tr><td>" + gWrokerKey[item] + "</td><td>" + gRoleDes[obj[item]] + "</td><tr>";
        } else if (item == KEY_CHECKINTIME) {
            var checkindate = new Date(obj[item] * 1000);
            worker_details_info += "<tr><td>" + gWrokerKey[item] + "</td><td>" 
                                 + checkindate.getFullYear() + "-" + (checkindate.getMonth() + 1)
                                 + "-" + checkindate.getDate() + "</td><tr>";
        } else if (item == KEY_CHECKOUTTIME) {
            var checkoutdate = new Date(obj[item] * 1000);
            worker_details_info += "<tr><td>" + gWrokerKey[item] + "</td><td>";
            if (obj[item] != 0) {
                worker_details_info += checkoutdate.getFullYear() + "-" + (checkoutdate.getMonth() + 1)
                                 + "-" + checkoutdate.getDate();
            }
            worker_details_info += "</td><tr>";
        } else if (item != KEY_PASSWORD && item != KEY_ID) {
            worker_details_info += "<tr><td>" + gWrokerKey[item] + "</td><td>" + obj[item] + "</td><tr>";
        }        
    }
    worker_details_info = '<table class="table table-bordered table-hover table-condensed bg-info"><tbody>'+
                         worker_details_info + '<tbody></table>';

    $("#myModalLabel").text("个人详情");
    $("#worker_detail").show().append(worker_details_info);//显示详情模态框内容
}

//删除用户
function delWorkerAction(data) {
    var workerid = $(data).parent().parent().children("td").first().text();
    var errcode = 1;

    console.log("workerid="+workerid);
    $.post(URL_WORKER,{Command: CMD_DEL_WORKER, Worker: workerid},function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });

        if (errcode == 1) {
            alert("删除失败");
        } else {
            alert("删除成功");
            updataWorkerList();
        }
    })

}

//显示Admin修改用户信息界面
function changeWorkerAction(data) {
    var index = data.parentNode.parentNode.rowIndex;
    var obj = workers[index-1];

    console.log("index = " + index);
    $("#WorkerId").val(obj.WorkerId).attr("readonly",true);
    $("#Password").val(obj.Password);
    $("#Name").val(obj.Name);
    $("#Sex").val(obj.Sex);
    $("#IdentifyNo").val(obj.IdentifyNo);
    $("#Title").val(obj.Title);
    $("#Comment").val(obj.Comment);

    $("#myModalLabel").text("个人信息修改");
    $("#registered_from").show();//显示详情模态框内容
    $("#modeupdatabtn").show();
}

//admin提交修改用户信息
function updataPersonInfo() {
    var workerId = $("#WorkerId").val();
    var password = $("#Password").val();
    var name = $("#Name").val();
    var sex = $("#Sex").val();
    var identifyNo = $("#IdentifyNo").val();
    var title = $("#Title").val();
    var comment = $("#Comment").val();
    var errcode = 1;

    $.get(URL_WORKER, {
        Command: CMD_UPD_WORKER, Worker: workerId, Password: password, Name: name, Sex: sex,
        IdentifyNo: identifyNo, Title: title, Comment: comment}, function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });

        if (errcode == 1) {
            alert("修改失败");
        } else {
            $("#myModal").modal("hide");
            alert("修改成功");
            updataWorkerList();
        }
    });
}

//显示修改个人密码界面
function changePersonPasswordUi(data) {
    var workerid=$(data).parent().parent().children("td").first().text();
    console.log("workerid"+workerid);

    $("#myid").val(workerid);
    $("#mypassword").val("");
    $("#myModalLabel").text("修改密码");
    $("#person_password").show();//显示详情模态框内容
    $("#modechangebtn").show();
}

//修改个人密码event
function changePersonPasswordEvent() {
    var workerid = $("#myid").val();
    var password = $("#mypassword").val();
    var errcode = 1;

    console.log("workerid"+workerid);
    console.log("password"+password);
    $.get(URL_WORKER, {
        Command: CMD_CHGPWD_WORKER, Worker: workerid, Password: password}, function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });
        if (errcode == 1) {
            alert("修改失败");
        } else {
            $("#myModal").modal("hide");
            alert("修改成功");
        }
    });
}

function updataWorkerList() {
    $("#userlist").empty();
    getWorkersInfo();//获取全部用户信息
}

//添加用户
function addWorker() {
    var workerId = $("#WorkerId").val();
    var password = $("#Password").val();
    var name = $("#Name").val();
    var sex = $("#Sex").val();
    var identifyNo = $("#IdentifyNo").val();
    var title = $("#Title").val();
    // checkInTime = $("#CheckInTime").val();
    // checkOutTime = $("#CheckOutTime").val();
    var comment = $("#Comment").val();
    var errcode = 1;

    console.log("workerId = " + workerId + ", password = " + password + ", name = " + name + ", title = " + title + ", identifyNo = ", + identifyNo);
    if (workerId != "" && password!= "" && name!= "" && title!="" && identifyNo != "") {
        $.get(URL_WORKER, {
                Command: CMD_ADD_WORKER, Worker: workerId, Password: password, Name: name, Sex: sex,
                IdentifyNo: identifyNo, Title: title, Comment: comment}, function (data) {
            $.each(data, function(key,value){
                if (key == KEY_ERRCODE) {
                    errcode = value;
                }
            });

            if (errcode == 1) {
                alert("用户已存在");
            } else {
                $("#myModal").modal("hide");
                alert("新增成功");
            }                
        });
    } else {
        alert("请输入正确的信息");
    }
}

//恢复初始密码
function resetPassword(data) {
    var errcode = 1;
    var workerid=$(data).parent().parent().children("td").first().text();
    //var password="888888";
    $.get(URL_WORKER, {Command: CMD_CHGPWD_WORKER, Worker:workerid, Password:STR_DEFAULT_PWD}, function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });

        if (errcode == 1) {
            alert("修改失败");
        } else {
            alert("修改成功");
        }
    });
}