<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>登录</title>

    <!-- Material Design fonts -->
    <!-- <link rel="stylesheet" href="http://fonts.googleapis.com/css?family=Roboto:300,400,500,700" type="text/css">
     <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
     <!-- Bootstrap -->
    <link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet">
    <!-- Bootstrap Material Design -->
    <link href="../static/dist/css/bootstrap-material-design.css" rel="stylesheet">
    <link href="../static/dist/css/ripples.min.css" rel="stylesheet">
</head>
<body>
<div class="app">
    <div class="container">
        <div class="row">
            <div class="col-md-4 col-md-offset-4 col-sm-6 col-sm-offset-3">
                <div class="alert alert-dismissible alert-danger" style="margin-bottom:0;">
                    <button type="button" class="close" data-dismiss="alert">×</button>
                    <strong>Oh snap!</strong>
                    <a href="javascript:void(0)" class="alert-link">Change a few things up</a> and try submitting again.
                </div>
                <div style="height:25px;"></div>
                <form class="form-horizontal well" id="login-form" action="#" method="post">
                    <fieldset>
                        <legend style="text-align:center">登录</legend>
                        <<< .xsrf_token >>>
                        <div class="form-group">
                            <label for="inputEmail" class="col-md-3 control-label">用户名</label>

                            <div class="col-md-9">
                                <input type="text" class="form-control" id="inputEmail" name="username" placeholder="手机号/邮箱">
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="inputPassword" class="col-md-3 control-label">密&emsp;码</label>

                            <div class="col-md-9">
                                <input type="password" class="form-control" id="inputPassword" name="password" placeholder="密码">
                            </div>
                        </div>

                        <div class="form-group">
                            <div class="col-md-offset-2 col-md-10">
                                <div class="togglebutton">
                                    <label>
                                        <input type="checkbox" checked="">记住密码
                                    </label>
                                </div>
                            </div>
                        </div>

                        <div class="form-group">
                            <div class="col-md-10 col-md-offset-2">
                                <a href="javascript:void(0)" id="submit" class="btn btn-raised btn-success">登录</a>
                            </div>
                        </div>
                    </fieldset>
                </form>
                <div class="well col-md-12">
                    忘记密码?<a href="forget">重置密码</a>
                </div>
            </div>
        </div>
    </div>
</div>
<!-- jQuery -->
<script src="//code.jquery.com/jquery-1.10.2.min.js"></script>
<!-- Twitter Bootstrap -->
<script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.6/js/bootstrap.min.js"></script>

<!-- Material Design for Bootstrap -->
<script src="../static/dist/js/material.min.js"></script>
<script src="../static/dist/js/ripples.min.js"></script>
<script src="../static/js/form-validate.js"></script>
<script>
    $.material.init({validate:false});
    $.forms.init(
     {
      form:"#login-form",
      submitBtn:"#submit",
      fields: [{
          field:"#inputEmail",
          rules:{
              required:{
                  submitOnly:false,
                  message:"用户名不能为空!"
              }
          }
    },{
        field:"#inputPassword",
        rules:{
            required:{
                submitOnly:false,
                message:"密码不能为空!"
            },
            min:{
                count:6,
                message:"密码长度不能少于6"
            },
            max:{
                count:8,
                message:"密码长度不能超过8"
            }
        }
    }]
    }
    );

    var form = <<< .form_check >>>;
</script>
</body>
</html>