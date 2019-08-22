const datatablelang = {
        "sProcessing": "处理中...",
        "sLengthMenu": "每页 _MENU_ 项",
        "sZeroRecords": "没有匹配结果",
        "sInfo": "   当前显示第 _START_ 至 _END_ 项，共 _TOTAL_ 项。",
        "sInfoEmpty": "当前显示第 0 至 0 项，共 0 项",
        "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
        "sInfoPostFix": "",
        "sSearch": "搜索:",
        "sUrl": "",
        "sEmptyTable": "表中数据为空",
        "sLoadingRecords": "载入中...",
        "sInfoThousands": ",",
        "oPaginate": {
            "sFirst": "首页",
            "sPrevious": "上页",
            "sNext": "下页",
            "sLast": "末页",
            "sJump": "跳转"
        },
        "oAria": {
            "sSortAscending": ": 以升序排列此列",
            "sSortDescending": ": 以降序排列此列"
        }
};

function selectTr(selector) {
    selector.on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            oTable.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
}

function arrangeData(arrData) {
    let pageIndex = 1;
    let pageSize = 10;
    let orderBy = "";
    let orderIndex = 0;
    let sort = "desc";
    for (i = 0;i<arrData.length;i++){
        if (arrData[i].name =="iDisplayStart"){
            pageIndex = arrData[i].value;
        }

        if (arrData[i].name =="iDisplayLength"){
            pageSize = arrData[i].value;
        }
         //排序方式
        if (arrData[i].name =="sSortDir_0"){
            sort = arrData[i].value;
        }
    }

    for (i = 0;i<arrData.length;i++){
        if (arrData[i].name =="mDataProp_"+orderIndex){
            orderBy = arrData[i].value;
        }
    }

    let newData =[];
    newData.push({name:"PageSize",value:pageSize});
    newData.push({name:"PageIndex",value:pageIndex});
    newData.push({name:"OrderBy",value:orderBy});
    newData.push({name:"Isdesc",value:sort =="desc" ? true : false});
    newData.push({name:"TotalItemCount",value:""});
    //newData.push({name:"OrderByWithAcsOrDesc",value:orderBy +" "+sort});

    return newData;
}