package commond

import (
	"errors"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	// 根命令
	Use: "git",
	// 短描述
	Short: "Git is a distributed version control system.",
	// 长描述
	Long: `Git is a free and open source distributed version control system
designed to handle everything from small to very large projects 
with speed and efficiency.`,
	// 实际执行
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	RootCmd.Execute()
}
