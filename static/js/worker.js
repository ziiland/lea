var workers = new Array();
// var workerId="";
// var password="";
// var name="";
// var sex="";
// var identifyNo="";
// var title="";
// var comment="";

$(document).ready(function(){
    getWorkersInfo();
    $("#commit_worker").click(function(){
        addWorker();
    });
});

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
        if((item != "Password")&&(item != "Id")) {
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
                Command: CMD_ADD_WORKER, WorkerId: workerId, Password: password, Name: name, Sex: sex,
                IdentifyNo: identifyNo, Title: title, Comment: comment
            }, function (data) {
                $.each(data, function(key,value){
                    if (key == "Errcode") {
                        errcode = value;
                    }
                });

                if (errcode == 1) {
                    alert("用户已存在");
                } else {
                    $("#registered").modal("hide");
                    alert("新增成功");
                }                
        });
    } else {
        alert("请输入正确的信息");
    }
}