package user

import (
	"errors"

	"github.com/ManiMuridi/gocmder/command"
)

func Create(user User) *createUserCmd {
	return &createUserCmd{user: user}
}

type createUserCmd struct {
	user User
}

func (c *createUserCmd) Exec() command.Result {
	if c.user.Name == "" {
		return command.Result{
			Error: errors.New("user name cannot be empty"),
			Data:  nil,
		}
	}

	return command.Result{
		Error: nil,
		Data:  c.user,
	}
}
