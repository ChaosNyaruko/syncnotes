package main

import (
	"flag"
)

// TODO: define my own var to implement list values (e.g. -o dotfile dollop )
// check https://pkg.go.dev/flag#Value.
var obj = flag.String("o", "./tmp", "where you want to launch the sync process")

func main() {
	flag.Parse()
	// var objs  = []string{}
}
