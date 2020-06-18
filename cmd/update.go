package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cluster string
	service string
	desired int64
	force   bool
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update ECS Service.",
	Long:  `Update ECS Service.`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs := NewEcsService()
		ecs.ServiceUpdate(cluster, service, desired, force)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "ECS Cluster Name")
	updateCmd.Flags().StringVarP(&service, "service", "s", "", "ECS Service Name")
	updateCmd.Flags().Int64VarP(&desired, "desired", "d", 0, "Desired Count")
	updateCmd.Flags().BoolVarP(&force, "force", "f", false, "Force New deploy.")
}
