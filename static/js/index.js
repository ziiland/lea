$(document).ready(function(){
    $.when(getDataFromBackend()).done(function(){
        displayFooter();
        displayHeader();
    });
});

