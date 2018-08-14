package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the application version",
		Run: func(cmd *cobra.Command, args []string) {
			getVersion()
		},
	}

	return cmd
}

func getVersion() {
	versionNumber := "0.0.1"
	fmt.Printf("Version: %#v\n", versionNumber)
}
