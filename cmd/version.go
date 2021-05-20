package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Version Number.",
	Long:  `Print Version Number.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
