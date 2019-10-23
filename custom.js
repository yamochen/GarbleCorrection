$(".custom-file-input").on("change", function() {
    var fileName = $(this).val().split("\\").pop();
    $(this).siblings(".custom-file-label").addClass("selected").html(fileName);
});

function upload(){
    var fileData = $('#uploadFile').prop('files')[0];

    var reader = new FileReader();
    reader.onload = function(readerEvt) {
        var binaryString = btoa(readerEvt.target.result);
        post(binaryString);
    };
    reader.readAsBinaryString(fileData);
}

function post(data){
    $.ajax({
        url: "http://127.0.0.1:8080/simplifiedGarbled/",
        type: "POST",
        data: JSON.stringify({"event":{"Content": data}}),
        dataType:'json',
        headers: {
            'Content-Type': 'application/json'
        },
        success: function (response) {
            console.log(response);
        },
        error: function(error){
            console.log("Something went wrong", error);
        }
    });
}