package auth
const(
	Login = 2   //if Login included,indicate that we must visit this page after login
	JumpBack = 4  //if JumpBack included,indicate that: if we have not login,
	              //it will redirect to login page with next value,like hostname/account/signin?next=account/info
)