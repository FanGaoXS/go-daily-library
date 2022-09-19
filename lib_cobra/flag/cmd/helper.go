package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Error(cmd *cobra.Command, args []string, err error) {
	log.Fatalf("execute command: %v args: %v error: %v", cmd.Name(), args, err)
}
