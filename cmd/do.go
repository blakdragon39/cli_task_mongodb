package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task/taskdb"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Marks a task as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument", arg)
				return
			}

			ids = append(ids, id)

			tasks, err := taskdb.AllTasks()
			if err != nil {
				fmt.Println("Something went wrong:", err.Error())
				return
			}

			for _, id := range ids {
				if id <= 0 || id > len(tasks) {
					fmt.Println("Invalid task number:", id)
					continue
				}

				task := tasks[id - 1]
				err := taskdb.DeleteTask(task.Id)
				if err != nil {
					fmt.Printf("Failed to complete task %d. Error: %s\n", id, err.Error())
				} else {
					fmt.Printf("Completed task %d\n", id)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}