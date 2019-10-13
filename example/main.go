package main

import (
	"github.com/ManiMuridi/goexec/command"
)

type cmd struct {
	result *command.Result
}

func (c *cmd) Result() *command.Result {
	return c.result
}

func (c *cmd) Exec() {
	//c.result = &command.Result{
	//	Error: nil,
	//	Data:  nil,
	//}
	//
	//c.result.Error = errors.New("A new error")
	//
	//c.result.Data = struct{ Name string }{"A Name"}
	//fmt.Println("Running")
}

func main() {
	//c := &cmd{}
	//command.Exec(c)
	//fmt.Println(c.Result())
}
