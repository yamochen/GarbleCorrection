var fixData = "";

$(".custom-file-input").on("change", function() {
    var fileName = $(this).val().split("\\").pop();
    $(this).siblings(".custom-file-label").addClass("selected").html(fileName);
    $("#progressbar").width("10%");
});

function upload(){
    $("#progressbar").width("30%");
    var fileData = $('#uploadFile').prop('files')[0];

    var reader = new FileReader();
    reader.onload = function(readerEvt) {
        var binaryString = btoa(readerEvt.target.result);
        post(binaryString);
    };
    reader.readAsBinaryString(fileData);
}

function post(data){
    $("#progressbar").width("50%");
    $.post( "http://127.0.0.1:8080/simplifiedGarbled", JSON.stringify({"Content": data}), 
        function (data) {
            $("#progressbar").width("70%");
            text = $.parseJSON(data); 
            fixData = text["Content"];

            var data = new Blob([fixData], {type: 'text/plain'});
            var url = window.URL.createObjectURL(data);

            $("#downloadFile").attr("href", url);
            $('#downloadFile').prop("disabled", false);

            $("#progressbar").width("100%");
        }
    );
}
