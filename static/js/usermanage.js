var number="1";
var userid="0001";
var username="willamzhang";
var rank="1";
var money="10000";
var creattime="2016-10-19";
var details= "1";

function appenduserlist() {
  return "<tr>"+"<td>1</td>"+"<td>0001</td>"+"<td>willamzhang</td>"+"<td>1</td>"+"<td>10000</td>"+"<td>2016-10-19</td>"+"<td>1</td>"+"</tr>";
}

$(function(){
  $("#userlist").append(appenduserlist());
});

$(document).ready(function(){
    $("#add").click(function(){$("#userlist").append(appenduserlist());})
});