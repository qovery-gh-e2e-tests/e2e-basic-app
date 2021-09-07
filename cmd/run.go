package cmd

import (
	"github.com/Qovery/e2e-app/pkg"
	// "fmt"
)	

func Execute(port string) {
	routes := []pkg.Route{
		{"test", "get"},
	}

	var toto  []string
	// fmt.Println(toto[2])

	pkg.CreateServer(routes, port)
}
