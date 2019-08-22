
//弹出错误提示的登录框
$.showErr = function(str, func) {
    // 调用show方法
    BootstrapDialog.show({
        type : BootstrapDialog.TYPE_DANGER,
        title : '错误 ',
        message : str,
        size : BootstrapDialog.SIZE_SMALL,//size为小，默认的对话框比较宽
        buttons : [ {// 设置关闭按钮
            label : '关闭',
            action : function(dialogItself) {
                dialogItself.close();
            }
        } ],
        // 对话框关闭时带入callback方法
        onhide : func
    });
};

$.showConfirm = function(str, funcok, funcclose) {
    BootstrapDialog.confirm({
        title : '确认',
        message : str,
        type : BootstrapDialog.TYPE_WARNING, // <-- Default value is
        // BootstrapDialog.TYPE_PRIMARY
        closable : true, // <-- Default value is false，点击对话框以外的页面内容可关闭
        draggable : true, // <-- Default value is false，可拖拽
        btnCancelLabel : '取消', // <-- Default value is 'Cancel',
        btnOKLabel : '确定', // <-- Default value is 'OK',
        btnOKClass : 'btn-warning', // <-- If you didn't specify it, dialog type
        size : BootstrapDialog.SIZE_SMALL,
        // 对话框关闭的时候执行方法
        onhide : funcclose,
        callback : function(result) {
            // 点击确定按钮时，result为true
            if (result) {
                // 执行方法
                funcok.call();
            }
        }
    });
};

$.showSuccessTimeout = function(str, func) {
    BootstrapDialog.show({
        type : BootstrapDialog.TYPE_SUCCESS,
        title : '成功 ',
        message : str,
        size : BootstrapDialog.SIZE_SMALL,
        buttons : [ {
            label : '确定',
            action : function(dialogItself) {
                dialogItself.close();
            }
        } ],
        // 指定时间内可自动关闭
        onshown : function(dialogRef) {
            setTimeout(function() {
                dialogRef.close();
            }, YUNM._set.timeout);
        },
        onhide : func
    });
};


$.showOpen = function (title,htmlName,ajaxfunc) {
  let th = BootstrapDialog.show({
        type : BootstrapDialog.TYPE_PRIMARY,
        title : title,
        message: $('<div></div>').load(htmlName),
        closable : false, //点击对话框以外的页面内容可关闭
        buttons: [{
            icon: 'glyphicon glyphicon-send',
            label: '确定',
            cssClass: 'btn-primary',
            autospin: true,
            action: ajaxfunc
        }, {
            label: '取消',
            action: function(dialogRef){
                dialogRef.close();
            }
        }]
    });
}