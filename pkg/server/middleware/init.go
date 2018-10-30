package middleware

//IMiddleware is an interface for struct middleware
type IMiddleware interface {
}

//Middleware is a struct for middleware web
type Middleware struct {
}

//InitMiddleware is a function to instantiate middleware
func InitMiddleware() *Middleware {
	return &Middleware{}
}
