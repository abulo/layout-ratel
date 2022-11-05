layui.define(['jquery'], function (exports) {
    var $ = layui.jquery;
    $('#layui-laypage-btn').on('click', function () {
        var result = {};
        var paramStr = decodeURI(window.document.location.search);
        var reg = /^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$/;
        var regExp = new RegExp(reg);
        if (paramStr) {
        	paramStr = paramStr.substring(1);
        	var params = paramStr.split('&');
            for (var p = 0; p < params.length; p++) {
                str = unescape(params[p].split('=')[1])
                if(regExp.test(str)){
                    str = str.replace("+", " ");
                }
            	result[params[p].split('=')[0]] = str;
			}
        }
        result["page"] = $('#page_num').val();
        result["per_num"] = $('#per_num').val();
        window.location.href = window.location.pathname + "?" + $.param(result)
    });
    exports('page');
});