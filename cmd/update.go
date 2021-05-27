package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update ECS Service.",
	Long:  `Update ECS Service.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" {
			fmt.Println("Please input `cluster` name.")
			os.Exit(1)
		}
		if service == "" {
			fmt.Println("Please input `service` name.")
			os.Exit(1)
		}
		ecs := NewEcsService()
		result := ecs.ServiceUpdate(cluster, service, taskdef, desired, force)
		if !result {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "ECS Cluster Name")
	updateCmd.Flags().StringVarP(&service, "service", "s", "", "ECS Service Name")
	updateCmd.Flags().StringVarP(&taskdef, "taskdef", "t", "", "ECS Task Definition Name")
	updateCmd.Flags().Int64VarP(&desired, "desired", "d", 0, "Desired Count")
	updateCmd.Flags().BoolVarP(&force, "force", "f", false, "Force New deploy.")
}
