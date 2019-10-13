package command

import "fmt"

type NotFound struct {
	Name string
	Desc string
}

func (n *NotFound) Exec() {
	if n.Desc == "" {
		n.Desc = fmt.Sprintf("Could not find command: %s", n.Name)
	}

	fmt.Println(n.Desc)
}
