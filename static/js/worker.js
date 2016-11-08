var workers = new Array();
$(document).ready(function(){
    getDataFromBackend();
    getWorkersInfo();
    setModalEvent();
    $(document).on(EVT_HIDE_BTNADD, function(){
        console.log("EVT_HIDE_BTNADD: title=" + title);
        //��ʾ����û���ť
        if (title == "Admin") {
        $("#add").show();
       } 
    });
});
//��ģ̬���¼�
function setModalEvent(){
    $("#registered").on("hidden.bs.modal", function() {
        $("#registered_from").reset;
    });
}

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

function descriptionWorkers(workers) {
    var worker_info = "";
    for ( item in workers) {
        console.log("item =" + item + ", value=" + workers[item]);
        if (item == KEY_TITLE) {
            worker_info = worker_info + "<td>" + gRoleDes[workers[item]] + "</td>";
        } else if ((item != KEY_PASSWORD) && (item != "Id")) {
            worker_info = worker_info + "<td>" + workers[item] + "</td>";
        }
    }
    $("#userlist").append("<tr>"+worker_info+"</tr>");
}

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
                $("#registered").modal("hide");
                alert("�����ɹ�");
            }                
        });
    } else {
        alert("��������ȷ����Ϣ");
    }
}