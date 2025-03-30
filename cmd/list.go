package cmd

import (
	"github.com/devamshi/tasks-cli/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTasks("todo")
	},
}
