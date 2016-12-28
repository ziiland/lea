var STR_CLEAN = "clean";
var CELL_LENGTH = 14;
var NFC_LENGTH = 16;

function CleanInput() {
    $("#CellId").val("");
    $("#NFCId").val("");
    $("#CellId").focus();
}

function ToggleFocus() {
    //toLowerCase
}

function SubmitInput(cell, nfc) {
  var errcode;
  var _farm = cell.slice(0, 6);
  var _cell = cell.slice(6);

  console.log("SubmitInput: farm=" + _farm + ", cell=" + _cell + ", nfc=" + nfc);
  $.get("/land/cell", {Command:"BindNFC", Farm:_farm, Cell:_cell, NFC:nfc}, function(data){
    console.log("data=" + data);

    $.each(data, function(key, value){
        if (key == "Errcode") {
            errcode = value;       
        }
    });
    console.log("Errcode=", errcode); 
    if (errcode == 1) {
        alert("该地块已经绑定其他NFC码，请检查确认");
    }
    //window.location.assign("./" + page);
  });

  CleanInput();
}

function HandleInput() {
    var cell = $("#CellId").val();
    var nfc = $("#NFCId").val();

    console.log("HandleInput: cell=" + cell + ", nfc=" + nfc);
    if (cell.length == 0) {
        // don't have cell id, focus this input item
        $("#NFCId").blur();
        $("#CellId").focus();
    } else if (nfc.length == 0) {
        // don't have nfc id, foucs this input item\
        $("#CellId").blur();
        $("#NFCId").focus();
    } else if (cell.length != CELL_LENGTH) {
        // cell id wrong
        $("#CellId").val("");
        $("#NFCId").blur();
        $("#CellId").focus();       
        alert("单元号输入有误！")
    } else if (nfc.length != NFC_LENGTH){
        $("#NFCId").val("");
        $("#CellId").blur();
        $("#NFCId").focus();
        alert("NFC号码输入有误！")
    } else {
        // submit the info
        SubmitInput(cell, nfc);      
    }
}

function Binding(){ 
    console.log("Binding!");
    $("#CellId").change(function(){
        var cell = $("#CellId").val().toLowerCase();
        console.log("Input CellId:" + cell);
        if (cell == STR_CLEAN) {
            CleanInput();
        } else if(cell != "") {
            HandleInput();
        } else {
            // do nothing
        }
    });

    $("#NFCId").change(function(){
        var nfc = $("#NFCId").val().toLowerCase();
          if (nfc == STR_CLEAN) {
              CleanInput();
          } else if(nfc != "") {
              HandleInput();
          } else {
              // do nothing
          }        
    });
}

$(document).ready(function(){
    console.log("READY!");
    $("#CellId").focus();
    Binding();
    // $(window).on("unload", function() {
    //     console.log("onfunc unload the window");
    //     //alert("unload the window");
    //     $.post("/mp/smt", {Command:"unload"}, function(data){
    //     });
    //   });             
});




// $(window).onbeforeunload = function() {
//   alert("onbeforeunload");
// };

// $(window).unload(function(){ 
//     alert("获取到了页面要关闭的事件了！"); 
// });
