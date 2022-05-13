const url = "http://localhost:1323";
const modalLogin = new bootstrap.Modal(document.getElementById('modalLogin'))
const modalTodo = new bootstrap.Modal(document.getElementById('modalTodo'))
const typeNotiSuccess = 'success';
const typeNotiError = 'error';
const typeNotiinfo = 'info';

function CallApi(url, method="GET", data={}, isAuth = fasle, callbackSuccess){
    var headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json' 
    }
    if(isAuth){
        headers["Authorization"] = "Bearer " + $.cookie('access_token');
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

function getTodos(){
    var urlApi = url + "/api/todos";
    CallApi(urlApi, 'GET', {}, true, (result)=>{
        todos = result.result ? result.result : [];
        const tbBody = $("#tableBodyTodos");
        tbBody.empty();
        todos.forEach(todo => {
            const tr = `<tr>
            <th scope="row">${todo.id}</th>
            <td>${todo.task}</td>
            <td>${todo.status === true ? '<i style="color:#02a518;" class="fa fa-check" aria-hidden="true"></i>' : '<i class="fa fa-clock-o" style="color:#a53d02;" aria-hidden="true"></i>'}</td>
            <td>
                <a type="button" class="btn-icon mr-2" id="btnUpdateTodo" title="Update Todo"><i class="fa fa-pencil fa-lg" aria-hidden="true"></i></a>
                <a type="button" class="btn-icon mx-2" id="btnDeleteTodo" title="Delete Todo"><i class="fa fa-trash fa-lg" aria-hidden="true"></i></a>
            </td>
            </tr>`;
            tbBody.append(tr);
        });
    })
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
