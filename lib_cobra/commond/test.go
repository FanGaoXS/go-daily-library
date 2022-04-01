package commond

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "version test",
	Run: func(cmd *cobra.Command, args []string) {
		// 执行命令
		output, err := ExecuteCommand("git", "version", args...)
		if err != nil {
			Error(cmd, args, err)
		}

		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	// 初始化时将testCmd加入versionCmd
	versionCmd.AddCommand(testCmd)
}
