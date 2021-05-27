package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List ECS Service in Cluster.",
	Long:  `List ECS Service in Cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" {
			fmt.Println("Please input `cluster` name.")
			os.Exit(1)
		}
		ecs := NewEcsService()
		result := ecs.Services(cluster)
		if !result {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)
	servicesCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "ECS Cluster Name")
}
