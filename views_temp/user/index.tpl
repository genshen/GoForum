<html lang="en" >
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
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

<div ng-controller="AppCtrl" layout="column" ng-cloak>
    <md-content layout-padding>
        <form name="projectForm">
            <md-input-container class="md-block">
                <label>Description</label>
                <input md-maxlength="30" required md-no-asterisk name="description" ng-model="project.description">
                <div ng-messages="projectForm.description.$error">
                    <div ng-message="required">This is required.</div>
                    <div ng-message="md-maxlength">The name has to be less than 30 characters long.</div>
                </div>
            </md-input-container>
            <md-input-container class="md-block">
                <label>Client Name</label>
                <input required name="clientName" ng-model="project.clientName">
                <div ng-messages="projectForm.clientName.$error">
                    <div ng-message="required">This is required.</div>
                </div>
            </md-input-container>
            <md-input-container class="md-block">
                <label>Client Email</label>
                <input required type="email" name="clientEmail" ng-model="project.clientEmail"
                       minlength="10" maxlength="100" ng-pattern="/^.+@.+\..+$/" />
                <div ng-messages="projectForm.clientEmail.$error" role="alert">
                    <div ng-message-exp="['required', 'minlength', 'maxlength', 'pattern']">
                        Your email must be between 10 and 100 characters long and look like an e-mail address.
                    </div>
                </div>
            </md-input-container>
            <md-input-container class="md-block">
                <label>Hourly Rate (USD)</label>
                <input required type="number" step="any" name="rate" ng-model="project.rate" min="800"
                       max="4999" ng-pattern="/^1234$/" />
                <div ng-messages="projectForm.rate.$error" multiple md-auto-hide="false">
                    <div ng-message="required">
                        You've got to charge something! You can't just <b>give away</b> a Missile Defense
                        System.
                    </div>
                    <div ng-message="min">
                        You should charge at least $800 an hour. This job is a big deal... if you mess up,
                        everyone dies!
                    </div>
                    <div ng-message="pattern">
                        You should charge exactly $1,234.
                    </div>
                    <div ng-message="max">
                        {{projectForm.rate.$viewValue | currency:"$":0}} an hour? That's a little ridiculous. I
                        doubt even Bill Clinton could afford that.
                    </div>
                </div>
            </md-input-container>
            <p style="font-size:.8em; width: 100%; text-align: center;">
                Make sure to include <a href="https://docs.angularjs.org/api/ngMessages" target="_blank">ngMessages</a> module when using ng-message markup.
            </p>
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
                    description: ' Missile Defense System',
                    rate: 500
                };
            });
</script>

</body>
</html>