var WorkerId="001";
var Name="aaa";
var Sex="男";
var IdentifyNo="3333333333333";
var Title="经理";
var CheckInTime="20161213";
var CheckOutTime="在职";
var Comment="dsafss";

function appenduserlist() {
  return "<tr><td>"+WorkerId+"</td><td>"+Name+"</td><td>"+Sex+"</td><td>"+
      IdentifyNo+"</td><td>"+Title+"</td><td>"+CheckInTime+"</td><td>"+CheckOutTime+"</td><td>"+Comment+"</td></tr>";
}

$(function(){
  $("#userlist").append(appenduserlist());
});

$(document).ready(function(){
    $("#add").click(function(){$("#userlist").append(appenduserlist());})
});