package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Id   string
	Uuid string
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove user",
	Long:  "remove user with given information",
	Run: func(cmd *cobra.Command, args []string) {
		removeUser(Id, Uuid)
	},
}

func removeUser(id string, uuid string) {
	// do remove
	if id != "" {
		fmt.Printf("remove user with id: %s\n", id)
	}
	if uuid != "" {
		fmt.Printf("remove user with uuid: %s\n", uuid)
	}
	if id != "" && uuid != "" {
		fmt.Printf("remove user with id: %s and uuid: %s\n", id, uuid)
	}
}

func init() {
	removeCmd.Flags().StringVar(&Id, "id", "", "id of user")
	removeCmd.Flags().StringVar(&Uuid, "uuid", "", "uuid of user")
	rootCmd.AddCommand(removeCmd)
}
