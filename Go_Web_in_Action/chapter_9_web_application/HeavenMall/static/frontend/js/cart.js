(function($){
    var app={
        init:function(){             
            this.changeCartNum();       
            this.deleteConfirm();
            this.initCheckBox();
            this.isCheckedAll();
        },
        deleteConfirm:function(){
            $('.delete').click(function(){    
                var flag=confirm('您确定要删除吗?');    
                return flag;    
            })
    
        },    
        initCheckBox(){
            //全选按钮点击
            $("#checkAll").click(function() {               
                if (this.checked) {
                    $(":checkbox").prop("checked", true);
                    $.get('/cart/changeAllCart?flag=1',function(response){
                        if(response.success){
                            $("#allPrice").html(response.allPrice+"元")                      
                        }
                    })
                  
                }else {
                    $(":checkbox").prop("checked", false);   
                    $.get('/cart/changeAllCart?flag=0',function(response){
                        if(response.success){
                            $("#allPrice").html(response.allPrice+"元")                      
                        }
                    })                    
                }

               
            });    

            // //点击单个选择框按钮的时候触发
            var _that=this;  //this是app对象            
            $(".cart_list :checkbox").click(function() {            
                _that.isCheckedAll(); 

                var product_id=$(this).attr('product_id');
                var product_color=$(this).attr('product_color');
                $.get('/cart/changeOneCart?product_id='+product_id+'&product_color='+product_color,function(response){

                    if(response.success){
                        $("#allPrice").html(response.allPrice+"元")                      
                    }
                })

            });
        },
        //判断全选是否选择
        isCheckedAll(){             
            var chknum = $(".cart_list :checkbox").size();//checkbox总个数
            var chk = 0;  //checkbox checked=true总个数
            $(".cart_list :checkbox").each(function () {  
                if($(this).prop("checked")==true){
                    chk++;
                }
            });
            if(chknum==chk){//全选
                $("#checkAll").prop("checked",true);
            }else{//不全选
                $("#checkAll").prop("checked",false);
            }
        }, 
        changeCartNum(){

            $('.decCart').click(function(){
                
                var product_id=$(this).attr('product_id');
                var product_color=$(this).attr('product_color');
          
                $.get('/cart/decCart?product_id='+product_id+'&product_color='+product_color,function(response){

                    if(response.success){
                        $("#allPrice").html(response.allPrice+"元")
                        //注意this指向
                        $(this).siblings(".input_center").find("input").val(response.num)
                        $(this).parent().parent().siblings(".totalPrice").html(response.currentAllPrice+"元")
                    }
                }.bind(this))

            });

            $('.incCart').click(function(){             
                var product_id=$(this).attr('product_id');
                var product_color=$(this).attr('product_color');
          
                $.get('/cart/incCart?product_id='+product_id+'&product_color='+product_color,function(response){

                    if(response.success){
                        $("#allPrice").html(response.allPrice+"元")
                        //注意this指向
                        $(this).siblings(".input_center").find("input").val(response.num)
                        $(this).parent().parent().siblings(".totalPrice").html(response.currentAllPrice+"元")
                    }
                }.bind(this))
               
            });
        }
    }

    $(function(){
        app.init();
    })    
})($)
