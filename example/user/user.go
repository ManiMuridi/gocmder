package user

import (
	"errors"
	"fmt"

	"github.com/ManiMuridi/gocmder/command"
)

func Create(user User) *createUserCmd {
	return &createUserCmd{user: user, result: &command.Result{}}
}

type createUserCmd struct {
	user   User
	result *command.Result
}

func (c *createUserCmd) Exec() {
	if c.user.Name == "" {
		c.result.Error = errors.New("user name cannot be empty")
		fmt.Println(fmt.Sprintf("Executing: %+v", c.result))
		return
	}

	c.result.Data = c.user
	fmt.Println(fmt.Sprintf("Executing: %+v", c.result))
}

func (c *createUserCmd) Result() *command.Result {
	return c.result
}
