package cmd

import (
	"github.com/Qovery/e2e-app/pkg"
	"fmt"
)	

func Execute(port string) {
	routes := []pkg.Route{
		{"test", "get"},
	}

	var toto  []string
	fmt.Println(toto)

	pkg.CreateServer(routes, port)
}


deploy 1 : OK

deploy 2 : failed

deploy 3: ok