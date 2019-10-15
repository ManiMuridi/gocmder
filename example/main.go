package main

import (
	"fmt"

	"github.com/ManiMuridi/gocmder"
	"github.com/ManiMuridi/gocmder/example/user"
)

func main() {
	createUser := user.Create(user.User{
		Name:  "",
		Email: "new.user@fakemail.com",
	})

	result := gocmder.Exec(createUser)

	if result.Error != nil {
		fmt.Println("Command failed to execute")
		fmt.Println(result.Error)
		return
	}

	fmt.Println("Command executed successfully!")

	//if createUser.result.Error != nil {
	//	fmt.Println(fmt.Sprintf("Command failed to execute: %s", createUser.result.Error))
	//} else {
	//	fmt.Println("Command executed successfully")
	//}

	//gocmder.Add(&SomeCmd{})
	//gocmder.ExecName("somecmd")
	//gocmder.Exec(&SomeCmd{})
}
