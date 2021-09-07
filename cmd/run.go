package cmd

import "github.com/Qovery/e2e-app/pkg"

func Execut(port string) {
	routes := []pkg.Route{
		{"test", "get"},
	}
	pkg.CreateServer(routes, port)
}
