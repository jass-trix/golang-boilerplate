package routes

//LoadRoute is a function that will instantiate all the route
func (r Router) LoadRoute() {
	r.Router.GET("/", r.Handler.HandleHelloWorld)
}
