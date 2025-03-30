package cmd

import (
	"fmt"

	"github.com/devamshi/tasks-cli/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Run: HandleAddTaskCmd,
}

func HandleAddTaskCmd(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("args cannot be empty")
		return
	}

	fmt.Println(task.AddTask(args[0]))
}
