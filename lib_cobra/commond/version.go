package commond

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show git version info.",
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
	// 初始化时将versionCmd加入rootCmd
	RootCmd.AddCommand(versionCmd)
}
