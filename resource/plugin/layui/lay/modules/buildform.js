layui.define(['jquery'], function(exports) {
	var $ = layui.jquery;

	window.buildForm = function(method,action,data)
	{

	    var form = $('<form></form>');
	    // 设置属性
	    form.attr('action', action);
		form.attr('method', method);
	    form.attr('enctype', 'multipart/form-data');
	    // 创建Input
	    for (var i = 0; i < data.length; i++) {
	        formInput = $('<input type="hidden" name="'+data[i].name+'" />');
	        formInput.attr('value', data[i].value);
	        form.append(formInput);
	    }
	    $(document.body).append(form);
	    // 提交表单
	    form.submit();
	    // console.log(form);
	    // 注意return false取消链接的默认动作
	    return false;
	}


	exports('buildform');
});
