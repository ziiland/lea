var prices = new Array();
///////////////////////////////////////////////////////////////////////////////////////
function PriceItemConstructor(name, kind, price, discount, show, comment) {
    this.Name = name;
    this.Kind = kind;
    this.Show = show;
    this.Price = price;
    this.Discount = discount;
    this.Comment = comment;
}

$(document).ready(function(){
    $.when(getDataFromBackend()).done(function(){
        LoadPriceInfo();
    });
    bindMyModalClick();
});
//绑定模态框关闭事件
function bindMyModalClick(){
    $("#myModal").on("hidden.bs.modal", function() {
        $("#price_from :input").val("");
        $("#modalSaveBtn").off("click");
    });
}
//重新加载全部价格信息
function reLoadPrices() {
    $("#price_list").empty();
    prices.length = 0;
    LoadPriceInfo();
}
//加载全部价格信息
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
//显示全部价格信息
function displayPriceTable() {
    var priceInfo = "";
    var changeBtn = '<button class="btn btn-info btn-sm" onclick="changeBtnAction(this)" data-toggle="modal" data-target="#myModal">修改</button>';
    var dellBtn = '<button class="btn btn-danger btn-sm" onclick="dellBtnAction(this)">删除</button>';
    console.log("prices=" + prices);
    $.each(prices,function (key,value) {
        for ( item in value) {
            priceInfo += "<td>" + value[item] + "</td>";
        }
        priceInfo =  "<tr>"+ priceInfo + "<td>"+ changeBtn + dellBtn+"</td>"+"</tr>";
    });

    $("#price_list").append(priceInfo);

}
//修改
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
    $("#myModalLabel").text("价格修改");
}
//增加
function addBtnAction() {
    $("#name").attr("readonly",false);
    $("#kind").attr("readonly",false);
    $("#myModalLabel").text("价格修改");
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

    var item = new PriceItemConstructor(name, kind, price, discount, "true", comment);
    var json = JSON.stringify(item);    
    $.post(URL_PRICE,{Command:CMD_UPDATE_PRICE, CmdPara:json},function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });
    }).done(function () {
        if (errcode == 1) {
            alert("修改失败");
        } else {
            reLoadPrices();
            $("#myModal").modal("hide");
            alert("修改成功");
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
    var errcode = "";

    var item = new PriceItemConstructor(name, kind, price, discount, "true", comment);
    var json = JSON.stringify(item);
    $.post(URL_PRICE,{Command: CMD_ADD_RRICE, CmdPara:json}, function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });
    }).done(function () {
        if (errcode == 1) {
            alert("保存失败");
        } else {
            reLoadPrices();
            $("#myModal").modal("hide");
            alert("保存成功");
        }
    });
}
//删除
function dellBtnAction(o) {
    var kind = $(o).parent().parent().children("td:eq(2)").text();
    var errcode ="";

    console.log("dellBtnAction: kind = " + kind);
    //根据Kind来删除
    $.post(URL_PRICE,{Command: CMD_DEL_PRICE, Kind:kind},function (data) {
        $.each(data, function(key,value){
            if (key == KEY_ERRCODE) {
                errcode = value;
            }
        });
    }).done(function () {
        if (errcode == 1) {
            alert("删除失败");
        } else {
            reLoadPrices();
            alert("删除成功");
        }
    });
}