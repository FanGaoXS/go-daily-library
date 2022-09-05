package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git",
	Short: "Git is a distributed version control system.",
	Long: `Git is a free and open source distributed version control system
designed to handle everything from small to very large projects 
with speed and efficiency.`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unimplemented command"))
	},
}

func Execute() {
	// 执行root及其子命令
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("execute root command failed")
	}
}
