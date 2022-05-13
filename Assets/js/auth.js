$( document ).ready(function() {
    $("#btnSignOut").on("click", () => { 
        const tbBody = $("#tableBodyTodos");
        tbBody.empty();
        $("#formTodo").trigger("reset");
        $("#formLogin").trigger("reset");
        $.removeCookie("access_token");
        $.removeCookie("refresh_token");
        $('#containerLogin').show();
        $('#containerTodo').hide();
    });

    $("#formLogin").on("submit", (e) => {
        e.preventDefault();
        const data = $("#formLogin").serializeArray();
        authParam = {};
        data.forEach(el => {
            
            if(["status"].includes(el.name)){
                authParam[el.name] = el.value === "on" ? true : false;
            }else{
                authParam[el.name] = el.value;
            }
        });
        var urlApi = url + `/api/auth/login`;
        CallApi(urlApi, "POST", authParam, true,  (result)=>{
            var auth = result.result ? result.result : [];
            console.log(auth)
            $("#formLogin").trigger("reset");
            $.cookie('access_token', auth.access_token);
            $.cookie('refresh_token', auth.refresh_token);
            $('#containerLogin').hide();
            $('#containerTodo').show();
            getTodos()
            showNoti(typeNotiSuccess, result.messages ? result.messages : "Success!!!")
        })
    });
});