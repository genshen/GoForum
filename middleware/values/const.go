package values

const (
	Login = 2 << iota    //if Login included,indicate that we must visit this page after login
	JumpBack     //if JumpBack included,indicate that: if we have not login,
	              //it will redirect to login page with next value,like hostname/account/signin?next=account/info
	LoginJSON    //if you are not in login state,will return a json data which indicate you are not login
)

//const for user status
const (
	FREEZING int = iota
	UNACTIVATED
	STATUS_ACTIVE
)

//message type
const (
	POST_REPLY = iota
	FOLLOW_ADD

)