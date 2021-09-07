package cmd

import "github.com/Qovery/e2e-app/pkg"

func Execute(port string) {
	routes := []pkg.Route{
		{"test", "get"},
	}

	var toto  []string
	print(toto)

	pkg.CreateServer(routes, port)
}


deploy 1 : OK

deploy 2 : failed

deploy 3: ok