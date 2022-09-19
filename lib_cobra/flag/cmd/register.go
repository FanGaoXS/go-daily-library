package cmd

import "github.com/spf13/cobra"

var (
	Sex  bool
	Name string
	Age  int
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register user",
	Long:  "register information of user",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	registerCmd.PersistentFlags().BoolVarP(&Sex, "sex", "s", false, "sex of user, default is boy(false)")
	registerCmd.Flags().StringVar(&Name, "name", "", "name of user")
	registerCmd.Flags().IntVarP(&Age, "age", "a", 0, "age of user")
	rootCmd.AddCommand(registerCmd)
}
