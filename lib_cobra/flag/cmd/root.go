package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "manage",
	Short: "management of user information",
	Long:  "management of user information",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unimplemented command"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("execute root command failed")
	}
}
