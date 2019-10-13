package main

import (
	"fmt"

	"github.com/ManiMuridi/gocmder/command"
)

type SomeCmd struct {
	result *command.Result
}

func (c *SomeCmd) Result() *command.Result {
	return c.result
}

func (c *SomeCmd) Exec() {
	fmt.Println("Executing some cmd")
}
