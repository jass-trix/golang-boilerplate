package auth

//IAuth is an interface for Auth struct
type IAuth interface{}

//Auth is a struct that will maintatin all the authentication
type Auth struct {
}

//InitAuth is a function to instantiate auth
func InitAuth() *Auth {
	return &Auth{}
}
