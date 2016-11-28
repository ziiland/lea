var workers = new Array();
var gWrokerKey = {"WorkerId":"工号", "Password":"密码", "Name":"姓名"
                , "Sex":"性别", "IdentifyNo":"身份证号", "Title":"角色"
                , "CheckInTime":"入职时间", "CheckOutTime":"离职时间", "Comment":"备注"};
///////////////////////////////////////////////////////////////////////////////
$(document).ready(function(){
    $.when(getDataFromBackend()).done(function(){
        getWorkersInfo();
        displayAddButton();
    });
    setModalEvent();//绑定模态框事件
});
//绑定模态框事件
function setModalEvent(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#worker_detail").empty().hide();
        $("#registered_from").hide();
        $("#registered_from :input").val("");
        $("#person_password").hide();
        $("#person_password :input").val("");
        $("#modesavebtn").hide().off("click");

    });
}
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
    $("#WorkerId").attr("readonly",false);
    $("#registered_from").show();
    $("#myModalLabel").text("新增员工");
    $("#modesavebtn").show().on("click",function () {
        addWorker();
    });
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
                updataWorkerList();
                $("#myModal").modal("hide");
                alert("新增成功");
            }
        });
    } else {
        alert("请输入正确的信息");
    }
}
//重新加载全部用户列表
function updataWorkerList() {
    workers.length = 0;
    $("#user_list").empty();
    getWorkersInfo();//获取全部用户信息
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
    var btn_state = "";
    var bgcolor ="";

    for ( item in workers) {
        switch (item){
            case KEY_WORKERID:
            case KEY_NAME:
            case KEY_SEX:
                worker_info += "<td>" +workers[item] + "</td>";
                break;
            case KEY_TITLE:
                worker_info += "<td>" +gRoleDes[workers[item]] + "</td>";
                break;
            case KEY_CHECKOUTTIME:
                if (workers[item] != 0){
                    bgcolor ="style='background-color: grey'"
                    btn_state = 'disabled="disabled"';
                }
                break;
            case KEY_CHECKINTIME:
                var time = timeToDate(workers[item]);
                worker_info += "<td>" + time + "</td>";
                break;
        }
    }
    console.log("gLoginInfo.title = "+ gLoginInfo.title)
    console.log("btn_state = " + btn_state);

    var detail_btn='<button class="btn btn-sm btn-info" onclick="workerDetailsAction(this)" data-toggle="modal" data-target="#myModal">详情</button>';
    var del_btn='<button class="btn btn-sm btn-danger"'+btn_state+'onclick="delWorkerAction(this)" >删除</button>';
    var change_btn='<button class="btn btn-sm btn-warning"'+btn_state+' onclick="changeWorkerAction(this)" data-toggle="modal" data-target="#myModal">修改</button>';
    var reset_btn='<button class="btn btn-sm btn-warning" onclick="resetPassword(this)" data-toggle="modal" data-target="#myModal">恢复密码</button>';
    var changePassword_btn='<button class="btn btn-sm btn-info"onclick="changePersonPasswordUi(this)" data-toggle="modal" data-target="#myModal">修改密码</button>';
    var showBtn ="";

    switch (gLoginInfo.title){
        case STR_ADMIN:
            if(workers[KEY_WORKERID] == STR_ADMIN){
                showBtn = detail_btn +"&nbsp;"+ change_btn;
            }else{
                showBtn = detail_btn +"&nbsp;"+ change_btn + "&nbsp;"+del_btn;
            }
            break;
        case STR_MANAGER:
            if (workers[KEY_WORKERID] == gLoginInfo.workerId.toUpperCase()) {
                showBtn = detail_btn + "&nbsp;" + changePassword_btn;
            }else{
                if(workers[KEY_TITLE] == STR_MANAGER){
                    showBtn = detail_btn;
                }else if(workers[KEY_TITLE] == STR_WORKER){
                    showBtn = detail_btn + "&nbsp;" +reset_btn;
                }
            }
            break;
        case STR_WORKER:
            showBtn = changePassword_btn;
            break;
    }
    worker_info += "<td>" + showBtn +"</td>";
    worker_info ="<tr "+bgcolor+">"+worker_info+"</tr>";

    $("#user_list").append(worker_info);
}

//显示用户详情
function workerDetailsAction(o) {
    var worker_details_info = "";
    var index = o.parentNode.parentNode.rowIndex;
    console.log("index=" + index);

    var obj = workers[index-1];
    for(item in obj){
        var title = ""
        var content = "";
        var show = true;
        switch (item){
            case KEY_TITLE:
                title = gWrokerKey[item];
                content = gRoleDes[obj[item]];
                show = true;
                break;
            case KEY_WORKERID:
            case KEY_NAME:
            case KEY_SEX:
            case KEY_IDENTIFYNO:
            case KEY_COMMENT:
                title = gWrokerKey[item];
                content = obj[item];
                show = true;
                break;
            case KEY_CHECKINTIME:
            case KEY_CHECKOUTTIME:
                title = gWrokerKey[item];
                content = timeToDate(obj[item]);
                show = true;
                break;
            default:
                show = false;
                break;
        }
        if(show){
            worker_details_info += "<tr><td>" + title+ "</td><td>" + content + "</td><tr>";
        }
    }
    worker_details_info = '<table class="table table-bordered table-hover table-condensed"><tbody>'+
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
    $("#modesavebtn").show().on("click",function () {
        updataPersonInfo();
    });
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
    $("#modesavebtn").show().on("click",function () {
        changePersonPasswordEvent()
    });
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