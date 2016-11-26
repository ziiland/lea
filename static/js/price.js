var prices = new Array();
///////////////////////////////////////////////////////////////////////////////////////
$(document).ready(function(){
    $.when(getDataFromBackend()).done(function(){
        LoadPriceInfo();
    });
    bindMyModalClick();
});
//��ģ̬��ر��¼�
function bindMyModalClick(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#price_from :input").val("");
        $("#modalSaveBtn").off("click");
    });
}
//���¼���ȫ���۸���Ϣ
function reLoadPrices() {
    $("#price_list").empty();
    prices.length = 0;
    LoadPriceInfo();
}
//����ȫ���۸���Ϣ
function LoadPriceInfo() {
    $.get(URL_PRICE, {Command:CMD_LOAD_PRICE},function (data) {
        $.each(data, function(key, value){
            console.log("key=" + key + ", value=" + value);
            if (key == KEY_PRICES) {
                $.each(value, function(index, obj){
                    prices[index] = obj;
                });
            }
        });
    }).done(function () {
        displayPriceTable();
    });
}
//��ʾȫ���۸���Ϣ
function displayPriceTable() {
    var priceInfo = "";
    var changeBtn = '<button class="btn btn-info btn-sm" onclick="changeBtnAction(this)" data-toggle="modal" data-target="#myModal">�޸�</button>';
    var dellBtn = '<button class="btn btn-danger btn-sm" onclick="dellBtnAction(this)">ɾ��</button>';
    console.log("prices=" + prices);
    $.each(prices,function (key,value) {
        for ( item in value) {
            priceInfo += "<td>" + value[item] + "</td>";
        }
        priceInfo =  "<tr>"+ priceInfo + "<td>"+ changeBtn + dellBtn+"</td>"+"</tr>";
    });

    $("#price_list").append(priceInfo);

}
//�޸�
function changeBtnAction(o) {
    var index = $(o).parent().parent().children("td:eq(0)").text();
    var priceInfo = prices[index-1];
    console.log("index="+index);
    $("#name").val(priceInfo.Name).attr("readonly",true);
    $("#kind").val(priceInfo.Kind).attr("readonly",true);
    $("#price").val(priceInfo.Price);
    $("#discount").val(priceInfo.Discount);
    $("#comment").val(priceInfo.Comment);

    $("#modalSaveBtn").on("click",function () {
        saveChangePrice();
    })
    $("#myModalLabel").text("�۸��޸�");
}
//����
function addBtnAction() {
    $("#name").attr("readonly",false);
    $("#kind").attr("readonly",false);
    $("#myModalLabel").text("�۸��޸�");
    $("#modalSaveBtn").on("click",function () {
        saveAddPrice();
    })
    $("#myModal").modal("show");
}
//save changeBtn
function saveChangePrice() {
    var name = $("#name").val();
    var kind = $("#kind").val();
    var price = $("#price").val();
    var discount = $("#discount").val();
    var comment = $("#comment").val();
    var errcode ="";
    $.post(URL_PRICE,{Command:CMD_UPDATE_PRICE,Name:name,Kind:kind,Price:price,Discount:discount,Comment:comment},function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });
    }).done(function () {
        if (errcode == 1) {
            alert("�޸�ʧ��");
        } else {
            reLoadPrices();
            $("#myModal").modal("hide");
            alert("�޸ĳɹ�");
        }
    });
}
//save addBtn
function saveAddPrice() {
    var name = $("#name").val();
    var kind = $("#kind").val();
    var price = $("#price").val();
    var discount = $("#discount").val();
    var comment = $("#comment").val();
    var errcode ="";

    $.post(URL_PRICE,{Command: CMD_ADD_RRICE,Name:name,Kind:kind,Price:price,Discount:discount,Comment:comment},function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });
    }).done(function () {
        if (errcode == 1) {
            alert("����ʧ��");
        } else {
            reLoadPrices();
            $("#myModal").modal("hide");
            alert("����ɹ�");
        }
    });
}
//ɾ��
function dellBtnAction(o) {
    var index = $(o).parent().parent().children("td:eq(0)").text();
    var errcode ="";
    //Ŀǰ�뷨�Ǹ���id��ɾ��
    $.post(URL_PRICE,{Command: CMD_DEL_PRICE,Id:index},function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });
    }).done(function () {
        if (errcode == 1) {
            alert("ɾ��ʧ��");
        } else {
            reLoadPrices();
            alert("ɾ���ɹ�");
        }
    });
}