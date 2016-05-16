<html lang="en" >
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>登录</title>
    <!-- Angular Material style sheet -->
    <link rel="stylesheet" href="http://ajax.googleapis.com/ajax/libs/angular_material/1.1.0-rc2/angular-material.min.css">
    <link rel="stylesheet" href="http://localhost:8080/static/css/main.css">
<style>
    .inputdemoErrors .inputErrorsApp {
        min-height: 48px; }
    .inputdemoErrors md-input-container > p {
        font-size: 0.8em;
        text-align: left;
        width: 100%; }
</style>
</head>
<body ng-app="BeeApp" ng-cloak>
<div ng-controller="AppCtrl" layout="column"  ng-cloak>
    <md-content layout-padding>
        <form name="projectForm" action="" method="post">
            <<<.xsrfdata>>>
            <md-input-container class="md-block">
                <label>用户名</label>
                <!--md-no-asterisk -->
                <input required name="username" ng-model="project.username">
                <div ng-messages="projectForm.username.$error" role="alert">
                    <div ng-message="error_password">用户名或密码错误</div>
                    <div ng-message="required">用户名不能为空</div>
                </div>
            </md-input-container>
            <md-input-container class="md-block">
                <label>密码</label>
                <input type="password" required name="password" ng-model="project.password"/>
                <div ng-messages="projectForm.password.$error" role="alert">
                    <div ng-message="error_password">用户名或密码错误</div>
                    <div ng-message="required">密码不能为空</div>
                </div>
            </md-input-container>
            <section layout="row" layout-align="center center" layout-wrap>
                <md-button type="submit" flex-xs="100" flex-sm="100"  class="md-raised md-primary">登&nbsp;录</md-button>
            </section>
        </form>
    </md-content>
</div>

<!-- Angular Material requires Angular.js Libraries -->
<script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.5.3/angular.min.js"></script>
<script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.5.3/angular-animate.min.js"></script>
<script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.5.3/angular-aria.min.js"></script>
<script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.5.3/angular-messages.min.js"></script>
<!-- Angular Material Library -->
<script src="http://ajax.googleapis.com/ajax/libs/angular_material/1.1.0-rc2/angular-material.min.js"></script>
<!-- Your application bootstrap  -->
<script type="text/javascript">
    angular.module('BeeApp', ['ngMaterial', 'ngMessages'])
            .controller('AppCtrl', function($scope) {
                $scope.project = {
                            username: <<<.username>>>
                };
            });
</script>

</body>
</html>