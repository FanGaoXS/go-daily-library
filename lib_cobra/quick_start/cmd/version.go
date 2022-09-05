package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show git version info.",
	Run: func(cmd *cobra.Command, args []string) {
		// 使用exec库直接执行 git version 命令
		command := exec.Command("git", "version")
		output, err := command.CombinedOutput()
		if err != nil {
			Error(cmd, args, err)
		}

		fmt.Print(string(output)) // 打印执行的结果
	},
}

func init() {
	// 将version命令添加为root命令的子命令
	rootCmd.AddCommand(versionCmd)
}
