var workers = new Array();

$(document).ready(function(){
    getWorkersInfo();
});

$(document).ready(function(){
    $("#commit_worker").click(function(){
        //addWorker
        $("#registered").modal("hide");
    });
});

function getWorkersInfo() {
        $.get(URL_WORKER, {Command:CMD_LOAD_WORKER}, function(data){
            $.each(data, function(key, value){
                if (key == KEY_WORKER) {
                    $.each(value, function(index, obj){
                        workers[index] = obj;
                        descriptionWorkers(obj);
                    });
                }
            });
        });
}

function descriptionWorkers(workers) {
    var worker_info ="";
    for ( item in workers) {
        console.log("item =" + item + ", value=" + workers[item]);
        if((item != "Password")&&(item != "Id")) {
            worker_info = worker_info + "<td>" + workers[item] + "</td>";
        }
    }
    $("#userlist").append("<tr>"+worker_info+"</tr>");
}

function addWorker() {

    var worker_form = {WorkerId:$("#WorkerId").val(),
                        Password:$("#Password").val(),
                        Name:$("#Name").val(),
                        Sex:$("#Sex").val(),
                        IdentifyNo:$("#IdentifyNo").val(),
                        Title:$("#Title").val(),
                        CheckInTime:$("#CheckInTime").val(),
                        CheckOutTime:$("#CheckOutTime").val(),
                        Comment:$("#Comment").val()};

    for ( item in worker_form) {
        console.log("item =" + item + ", value=" + worker_form[item]);
    }
}