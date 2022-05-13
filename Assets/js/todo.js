$( document ).ready(function() {
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
    $("#btnLoadTodos").on("click", ()=>{
        getTodos();
    });

    $("#btnAddTodo").on("click", () => { 
        $("#formTodo").trigger("reset");
        modalTodo.show();
    });

    $('.table').on('click', '#btnUpdateTodo', function() {
        var id = getID($(this).closest("tr"));
        var urlApi = url + `/api/todos/${id}`;
        CallApi(urlApi, 'GET', {}, true, (result)=>{
            var todo = result.result ? result.result : {};
            $("#todo_id").val(todo.id);
            $("#todo_task").val(todo.task);
            $('#todo_status').prop("checked", todo.status === true ? true : false);

        })
        modalTodo.show();
     });
    
    $('.table').on('click', '#btnDeleteTodo', function() {
        var $row = $(this).closest("tr");
        var id = getID($row);
        var cf = confirm("Are you sure for delete this task?");
        if(cf){
            var urlApi = url + `/api/todos/${id}`;
            CallApi(urlApi, 'DELETE', {}, true, (result)=>{
                $row.remove();
                showNoti(typeNotiSuccess , "Delelte todo success!!!");
            })
        }
    });

    $("#formTodo").on("submit", (e) => {
        e.preventDefault();
        const data = $("#formTodo").serializeArray();
        todo = {};
        todo["status"] = false;
        data.forEach(el => {
            
            if(["status"].includes(el.name)){
                todo[el.name] = el.value === "on" ? true : false;
            }else{
                todo[el.name] = el.value;
            }
        });
        id = +todo.id;
        delete todo.id;
        var method = (id > 0) ? "PATCH" : "POST";
        var urlApi = url + `/api/todos${(id > 0) ? "/"+id : ""}`;
        CallApi(urlApi, method, todo, true,  (result)=>{
            getTodos();
            var todo = result.result ? result.result : [];
            $("#formTodo").trigger("reset");
            modalTodo.hide();
            showNoti(typeNotiSuccess, result.messages ? result.messages : "Success!!!")
        })
    });

});