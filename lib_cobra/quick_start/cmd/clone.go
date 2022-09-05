package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone <url> [destination]",
	Short: "clone a repository into a new directory",
	Args: func(cmd *cobra.Command, args []string) error {
		// validate args
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		args = append([]string{"clone"}, args...) // 将 clone 添加到 args 前，以组成 git clone args
		command := exec.Command("git", args...)
		output, err := command.CombinedOutput()
		if err != nil {
			Error(cmd, args, err)
		}

		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
