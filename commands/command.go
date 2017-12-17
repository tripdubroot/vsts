package commands

import (
	"fmt"

	m "github.com/tripdubroot/vsts/models"
)

// Exec executes the command passed by args
func Exec(cmd string) {
	var result m.ProjectList
	result = execListProjects(cmd)

	for i := 1; i < result.Count; i++ {
		fmt.Println(result.Value[i].Name)
	}

	fmt.Println(cmd)
}