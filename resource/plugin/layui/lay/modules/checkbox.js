layui.define(['form','jquery'], function(exports) {
	var form = layui.form
 		$ = layui.jquery;

	form.on('checkbox(all)', function(data){
		var child = $(data.elem).parents('table').find('tbody input[type="checkbox"]');
			child.each(function(index, item){
				item.checked = data.elem.checked;
			});
		form.render('checkbox');
	});


	window.getCheckboxValue = function()
	{
		var adIds = "";
	    $("tbody > tr  input[type=checkbox]:checked").each(function(i) {
	        if (0 == i) {
	            adIds = $(this).val();
	        } else {
	            adIds += ("," + $(this).val());
	        }
	    });
	    return adIds;
	}

	// function getCheckboxValue() {
    //
	// }


	exports('checkbox');
});
