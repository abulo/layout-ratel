layui.define(['jquery'], function(exports) {
	var $ = layui.jquery;
	$(".content_style td").each(function(i, $el) {
		if ($el.innerHTML == "") {
			$el.innerHTML = "&nbsp;";
		}
	});
	tr_num = $(".content_style tr").length;
    td_num = $(".content_style tr").first().find("td").length;
    html = '';
    if (td_num == 0){
        td_num = $(".content_style").prev().first().find("th").length;
    }
    pre_num = $("#td_num").val()
    if (pre_num == undefined || pre_num == 0)
    {
        pre_num = 15;
    }
	for (i=0;i<pre_num-tr_num;i++)
    {
        html += '<tr>';
            for(j=0;j<td_num;j++)
            {
                html += '<td>&nbsp;</td>';
            }
        html += '</tr>';
    }
	$(html).appendTo(".content_style");
	exports('tablelist');
});
