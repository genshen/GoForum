<!DOCTYPE html>
<html lang="en">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Material Design fonts -->
    <link rel="stylesheet" href="http://fonts.googleapis.com/css?family=Roboto:300,400,500,700" type="text/css">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!-- Bootstrap -->
    <link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet">
    <!-- Bootstrap Material Design -->
    <link href="dist/css/bootstrap-material-design.css" rel="stylesheet">
    <link href="dist/css/ripples.min.css" rel="stylesheet">
</head>
<body>
<form class="form-horizontal">
    <fieldset>
        <legend>Legend</legend>
        <div class="form-group">
            <label for="inputEmail" class="col-md-2 control-label">Email</label>

            <div class="col-md-10">
                <input type="email" class="form-control" id="inputEmail" placeholder="Email">
            </div>
        </div>
        <div class="form-group">
            <label for="inputPassword" class="col-md-2 control-label">Password</label>

            <div class="col-md-10">
                <input type="password" class="form-control" id="inputPassword" placeholder="Password">
            </div>
        </div>

        <div class="form-group">
            <div class="col-md-offset-2 col-md-10">
                <div class="togglebutton">
                    <label>
                        <input type="checkbox" checked=""> Toggle button
                    </label>
                </div>
            </div>
        </div>

        <div class="form-group">
            <div class="col-md-10 col-md-offset-2">
                <button type="button" class="btn btn-default">Cancel</button>
                <button type="submit" class="btn btn-primary">Submit</button>
            </div>
        </div>
    </fieldset>
</form>
<!-- jQuery -->
<script src="//code.jquery.com/jquery-1.10.2.min.js"></script>
<!-- Twitter Bootstrap -->
<script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.6/js/bootstrap.min.js"></script>

<!-- Material Design for Bootstrap -->
<script src="dist/js/material.js"></script>
<script src="dist/js/ripples.min.js"></script>
<script>
    $.material.init();
</script>


</body>
</html>