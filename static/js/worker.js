var workers = new Array();
$(document).ready(function(){
    getDataFromBackend();
    displayFooter();//��ʾҳͷ
    displayHeader();//��ʾҳβ
    getWorkersInfo();//��ȡȫ���û���Ϣ
    setModalEvent();//��ģ̬���¼�


    console.log("loginInfo.title="+loginInfo.title)
    if (loginInfo.title== "Admin") {
        $("#add").show(); //��ʾ����û���ť
    }
});

//����û���ťevent
function addClickAction() {
    $("#WorkerId").attr("readonly",false);
    $("#registered_from").show();
    $("#myModalLabel").text("����Ա��");
    $("#modesavebtn").show();
}
//��ģ̬���¼�
function setModalEvent(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#worker_detail").empty().hide();
        $("#registered_from").hide().reset;
        $("#modesavebtn").hide();
        $("#modeupdatabtn").hide();
        $("#person_password").hide();//��ʾ����ģ̬������
        $("#modechangebtn").hide();
    });
}
//��ȡȫ���û���Ϣ
function getWorkersInfo() {
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

//��ʾ�û���Ϣtable
function descriptionWorkers(workers) {
    var worker_info = "";
    var btn_state='disabled="disabled"';
    for ( item in workers) {
        console.log("item =" + item + ", value=" + workers[item]);
        if (item == KEY_TITLE) {
            worker_info = worker_info + "<td>" + gRoleDes[workers[item]] + "</td>";
        } else if(item == KEY_CHECKOUTTIME){
            if(workers[item]==0){
                btn_state="";
            }
            worker_info = worker_info + "<td>" + workers[item] + "</td>";
        }else if ((item != KEY_PASSWORD) && (item != "Id")&&(item != "IdentifyNo")&&(item != "Comment")) {
            worker_info = worker_info + "<td>" + workers[item] + "</td>";
        }
    }
    console.log("loginInfo.title"+loginInfo.title)
    console.log("btn_state =" + btn_state);
    var detail_btn='<button class="btn btn-sm btn-info" onclick="workerDetailsAction(this)" data-toggle="modal" data-target="#myModal">����</button>';
    var del_btn='<button class="btn btn-sm btn-danger"'+btn_state+'onclick="delWorkerAction(this)" >ɾ��</button>';
    var change_btn='<button class="btn btn-sm btn-warning" onclick="changeWorkerAction(this)" data-toggle="modal" data-target="#myModal">�޸�</button>';
    var reset_btn='<button class="btn btn-sm btn-warning" onclick="resetPassword(this)" data-toggle="modal" data-target="#myModal">�޸�</button>';
    var changePassword_btn='<button class="btn btn-sm btn-info"onclick="changePersonPasswordUi(this)" data-toggle="modal" data-target="#myModal">�޸�����</button>';
    if(loginInfo.title== "Admin") {
       if(workers[KEY_TITLE]=="Admin") {
           worker_info = worker_info + "<td>" + detail_btn + change_btn + "</td>"
       }else{
           worker_info = worker_info + "<td>" + detail_btn +del_btn+ change_btn + "</td>"
       }
    }else if(loginInfo.title== "Manager"){
        worker_info = worker_info + "<td>" + detail_btn + reset_btn + "</td>"
    }
    else{
        worker_info = worker_info +"<td>" +changePassword_btn+"</td>";
    }
    $("#userlist").append("<tr>"+worker_info+"</tr>");
}

//��ʾ�û�����
function workerDetailsAction(o) {
    var worker_details_info = "";
    var index=o.parentNode.parentNode.rowIndex;
    console.log("index=" + index);
    var obj=workers[index-1];
    for(item in obj){
        worker_details_info = worker_details_info +"<tr><td>"+item+"</td><td>"+obj[item]+"</td><tr>";
    }
    worker_details_info= '<table class="table table-bordered table-hover table-condensed bg-info"><tbody>'+
                         worker_details_info+'<tbody></table>';

    $("#myModalLabel").text("��������");
    $("#worker_detail").show().append(worker_details_info);//��ʾ����ģ̬������

}

//ɾ���û�
function delWorkerAction(data) {
    var workerid=$(data).parent().parent().children("td").first().text();
    var errcode;
    console.log("workerid="+workerid);
    $.post(URL_WORKER,{Command: CMD_DEL_WORKER, Worker: workerid},function (data) {
        $.each(data, function(key,value){
            if (key == "Errcode") {
                errcode = value;
            }
        });

        if (errcode == 1) {
            alert("ɾ��ʧ��");
        } else {
            alert("ɾ���ɹ�");
            updataWorkerList();

        }
    })

}

//��ʾAdmin�޸��û���Ϣ����
function changeWorkerAction(data) {
    var index=data.parentNode.parentNode.rowIndex;
    console.log("index"+index);
    var obj=workers[index-1];
    $("#WorkerId").val(obj.WorkerId).attr("readonly",true);
    $("#Password").val(obj.Password);
    $("#Name").val(obj.Name);
    $("#Sex").val(obj.Sex);
    $("#IdentifyNo").val(obj.IdentifyNo);
    $("#Title").val(obj.Title);
    $("#Comment").val(obj.Comment);

    $("#myModalLabel").text("������Ϣ�޸�");
    $("#registered_from").show();//��ʾ����ģ̬������
    $("#modeupdatabtn").show();
}

//admin�ύ�޸��û���Ϣ
function updataPersonInfo() {
    var workerId = $("#WorkerId").val();
    var password = $("#Password").val();
    var name = $("#Name").val();
    var sex = $("#Sex").val();
    var identifyNo = $("#IdentifyNo").val();
    var title = $("#Title").val();
    var comment = $("#Comment").val();
    var errcode;

    $.get(URL_WORKER, {
        Command: CMD_UPD_WORKER, Worker: workerId, Password: password, Name: name, Sex: sex,
        IdentifyNo: identifyNo, Title: title, Comment: comment}, function (data) {
        $.each(data, function(key,value){
            if (key == "Errcode") {
                errcode = value;
            }
        });

        if (errcode == 1) {
            alert("�޸�ʧ��");
        } else {
            $("#myModal").modal("hide");
            alert("�޸ĳɹ�");
            updataWorkerList();
        }
    });
}
//��ʾ�޸ĸ����������
function changePersonPasswordUi(data) {
    var workerid=$(data).parent().parent().children("td").first().text();
    console.log("workerid"+workerid);

    $("#myid").val(workerid);
    $("#mypassword").val("");
    $("#myModalLabel").text("�޸�����");
    $("#person_password").show();//��ʾ����ģ̬������
    $("#modechangebtn").show();
}
//�޸ĸ�������event
function changePersonPasswordEvent() {
    var workerid = $("#myid").val();
    var password = $("#mypassword").val();
    console.log("workerid"+workerid);
    console.log("password"+password);
    $.get(URL_WORKER, {
        Command: CMD_CHGPWD_WORKER, Worker: workerid, Password: password}, function (data) {
        $.each(data, function(key,value){
            if (key == "Errcode") {
                errcode = value;
            }
        });
        if (errcode == 1) {
            alert("�޸�ʧ��");
        } else {
            $("#myModal").modal("hide");
            alert("�޸ĳɹ�");
        }
    });
}
function updataWorkerList() {
    $("#userlist").empty();
    getWorkersInfo();//��ȡȫ���û���Ϣ
}
//����û�
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

    console.log("workerId = " + workerId + ", password = " + password + ", name = " + name + ", title = " + title + ", identifyNo = ", + identifyNo);
    if (workerId != "" && password!= "" && name!="" && title!="" && identifyNo!="") {
        $.get(URL_WORKER, {
                Command: CMD_ADD_WORKER, Worker: workerId, Password: password, Name: name, Sex: sex,
                IdentifyNo: identifyNo, Title: title, Comment: comment}, function (data) {
            $.each(data, function(key,value){
                if (key == "Errcode") {
                    errcode = value;
                }
            });

            if (errcode == 1) {
                alert("�û��Ѵ���");
            } else {
                $("#myModal").modal("hide");
                alert("�����ɹ�");
            }                
        });
    } else {
        alert("��������ȷ����Ϣ");
    }
}
//�ָ���ʼ����
function resetPassword(data) {
    var workerid=$(data).parent().parent().children("td").first().text();
    var password="888888";
    $.get(URL_WORKER, {
        Command: CMD_CHGPWD_WORKER, Worker: workerid, Password: password}, function (data) {
        $.each(data, function(key,value){
            if (key == "Errcode") {
                errcode = value;
            }
        });
        if (errcode == 1) {
            alert("�޸�ʧ��");
        } else {
            alert("�޸ĳɹ�");
        }
    });
}