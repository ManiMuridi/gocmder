package main

import "github.com/ManiMuridi/gocmder"

func main() {
	gocmder.Add(&SomeCmd{})
	gocmder.ExecName("somecmd")
	gocmder.Exec(&SomeCmd{})
}
