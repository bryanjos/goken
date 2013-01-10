package main

import (
	"fmt"
	"github.com/paulbellamy/mango"
)

func hello(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 200, mango.Headers{}, mango.Body("Hello World!")
}

func bye(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 200, mango.Headers{}, mango.Body("Bye Bye!")
}

func routeNotFound(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 404, mango.Headers{}, mango.Body("You probably got lost :(")
}

func StartServer() {
	routes := make(map[string]mango.App)
	routes["/hello"] = new(mango.Stack).Compile(hello)
	routes["/bye"] = new(mango.Stack).Compile(bye)

	testServer := new(mango.Stack)
	testServer.Middleware(mango.ShowErrors("<html><body>{Error|html}</body></html>"), mango.Routing(routes))
	testServer.Address = "localhost:" + Configuration.Server_Port
	testServer.Run(routeNotFound)
	fmt.Printf("Running server on: %s\n", testServer.Address)
}
