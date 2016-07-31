package values

const (
	Login = 2   //if Login included,indicate that we must visit this page after login
	JumpBack = 4  //if JumpBack included,indicate that: if we have not login,
	//it will redirect to login page with next value,like hostname/account/signin?next=account/info
	LoginJSON = 8  //if you are not in login state,will return a json data which indicate you are not login
)

//const for user status
const (
	FREEZING int = 0
	UNACTIVATED int = 9
	STATUS_ACTIVE int = 10
)