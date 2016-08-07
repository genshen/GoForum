package utils

type Person struct {
	ID     uint
	Name   string
	Avatar string
}

//  </struct for me >
type UserStatus struct {
	Person
	IsLogin bool
}
//< /struct for me page >

type SimpleJsonResponse struct {
	Status   int
	Error    interface{}
	Addition interface{}
}
