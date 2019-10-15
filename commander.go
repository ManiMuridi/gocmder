package gocmder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/ManiMuridi/gocmder/command"
)

var (
	commands = make([]command.Command, 0)
)

type nilCmd struct {
	Name string
}

func (n *nilCmd) Result() *command.Result {
	return nil
}

func (n *nilCmd) Exec() command.Result {
	return command.Result{
		Error: errors.New(fmt.Sprintf("Could not find command: %s", n.Name)),
		Data:  nil,
	}
}

func Add(cmd command.Command) {
	commands = append(commands, cmd)
}

func FindCmd(cmdName string) command.Command {
	for i := range commands {
		c := commands[i]
		if strings.ToLower(reflect.TypeOf(c).Elem().Name()) == strings.ToLower(cmdName) {
			return c
		}
	}

	return &nilCmd{cmdName}
}

func ExecName(cmdName string) {
	cmd := FindCmd(cmdName)
	cmd.Exec()
}

func Exec(cmd command.Command) command.Result {
	return cmd.Exec()
}
