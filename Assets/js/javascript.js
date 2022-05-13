const url = "http://localhost:1323";
const modalLogin = new bootstrap.Modal(document.getElementById('modalLogin'))
const modalTodo = new bootstrap.Modal(document.getElementById('modalTodo'))
const typeNotiSuccess = 'success';
const typeNotiError = 'error';
const typeNotiinfo = 'info';

function CallApi(url, method="GET", data={}, isAuth = fasle, callbackSuccess){
    $.cookie('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkxlIFZhbiBUcm9uZyIsImVtYWlsIjoidHJvbmdsZXZhbjk4QGdtYWlsLmNvbSIsImV4cCI6MTY1MjQzNTU2N30.K_5zhos-_VnngQcKCVFVRaD_rPt5cJljAN9Qt7I-B3s');
    var headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json' 
    }
    if(isAuth){
        headers["Authorization"] = "Bearer " + $.cookie('token');
    }
    $.ajax({
        type: method,
        url: url,
        headers: headers,
        data: JSON.stringify(data),
        dataType: "json",
        contentType: "Application/json",
        beforeSend: function() {
            $(".loading").show();
        },
        success: function (result, status, xhr) {
            $(".loading").hide();
            callbackSuccess(result);
        },
        error: function (xhr, status, error) {
            $(".loading").hide();
            var res = xhr.responseJSON;
            if(res && res.errors){
                console.log(res.errors);
                messages = "";
                res.errors.forEach(element => {
                    messages+=element.message + "\n"
                });
                showNoti(typeNotiError, messages);
            }else{
                showNoti(typeNotiError, xhr.status + " " + xhr.statusText)
            }
        }
    });
}
function getID($row){
    var id = 0;
    $row.find("th:nth-child(1)").each(function() {
        id = $(this).text();
    });
    return id;
}

function showNoti(type = "success", message){
    toastr.options = {
    "closeButton": true,
    "newestOnTop": false,
    "progressBar": true,
    "positionClass": "toast-top-right",
    "preventDuplicates": false,
    "onclick": null,
    "showDuration": "300",
    "hideDuration": "1000",
    "timeOut": "5000",
    "extendedTimeOut": "1000",
    "showEasing": "swing",
    "hideEasing": "linear",
    "showMethod": "fadeIn",
    "hideMethod": "fadeOut"
    }
    switch (type) {
        case 'success':
            toastr.success(message);
            break;
        case 'error':
            toastr.error(message);
            break;
        default:
            toastr.info(message);
            break;
    }
}
