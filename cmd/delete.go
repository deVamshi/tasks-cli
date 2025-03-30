package cmd

import (
	"fmt"
	"strconv"

	"github.com/devamshi/tasks-cli/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Run: HandleDeleteTaskCmd,
}

func HandleDeleteTaskCmd(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("id is required to delete")
		return
	}

	idInt, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("ERROR: provide a valid id")
		fmt.Println(err.Error())
	}

	task.Delete(idInt)
}
