//config.adminPath已在base.js中已定义，当前页面可以直接使用，引入js的时需要注意先引入base.js，然后再引入当前js
$(function(){		
	productApp.init()
})
var productApp={
	init:function(){
        this.initFroalaEditor();
        this.initProductType();
        this.initPhotoUploader();
        this.initRelationProduct();
        this.initDeleteProductImage()
    },	
     //配置富文本编辑器
	initFroalaEditor:function(){
        new FroalaEditor('#content', {
            height: 200,
            language: 'zh_cn',
            imageUploadURL: '/'+config.adminPath+'/product/goUpload'
        });
    },
     //动态生成商品规格参数
	initProductType:function(){
		$("#product_type_id").change(function () {
            var cate_id = $(this).val()
            var str = ""
            var data = ""
            $.get('/'+config.adminPath+'/product/getProductTypeAttribute', { "cate_id": cate_id }, function (response) {
                console.log(response)
                if (response.success) {
                    data = response.result;
                    for (var i = 0; i < data.length; i++) {
                        if (data[i].attr_type == 1) {
                            str += '<li><span>' + data[i].title + ': 　</span>  <input type="hidden" name="attr_id_list" value="' + data[i].id + '" />   <input type="text" name="attr_value_list" /></li>'
                        } else if (data[i].attr_type == 2) {
                            str += '<li><span>' + data[i].title + ': 　</span> <input type="hidden" name="attr_id_list" value="' + data[i].id + '">  <textarea cols="50" rows="3" name="attr_value_list"></textarea></li>'
                        } else {
                            var attrArray = data[i].attr_value.split("\n")
                            str += '<li><span>' + data[i].title + ': 　</span>  <input type="hidden" name="attr_id_list" value="' + data[i].id + '" />';
                            str += '<select name="attr_value_list">'
                            for (var j = 0; j < attrArray.length; j++) {
                                str += '<option value="' + attrArray[j] + '">' + attrArray[j] + '</option>';
                            }
                            str += '</select>'
                            str += '</li>'
                        }
                    }
                    $("#product_type_attribute").html(str);

                }
            })
        })
	},
    //批量上传图片
    initPhotoUploader(){
        $('#photoUploader').diyUpload({
            url: '/'+config.adminPath+'/product/goUpload',
            success: function (response) {
                console.info(response);
                var photoStr = '<input type="hidden" name="product_image_list" value=' + response.link + ' />';
                $("#photoList").append(photoStr)
            },
            error: function (err) {
                console.info(err);
            }
        });

    },
    //修改颜色
    initRelationProduct(){
        $(".relation_product_color").change(function(){
            var color_id=$(this).val();
            var product_image_id=$(this).attr("product_image_id");
            $.get('/'+config.adminPath+'/product/changeProductImageColor',{color_id:color_id,product_image_id:product_image_id},function(response){
                    console.log(response);
            });
        })
    },
    //删除图库信息
    initDeleteProductImage(){
        $(".product_image_delete").click(function(){
            var product_image_id=$(this).attr("product_image_id");
            var flag = confirm("确定要删除吗?");
            var _that=this;
            if(flag){
                $.get('/'+config.adminPath+'/product/removeProductImage',{product_image_id:product_image_id},function(response){
                    if(response.success){
                        $(_that).parent().remove()
                    }
               });

            }
        })
    }

}
