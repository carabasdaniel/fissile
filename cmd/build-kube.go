package cmd

import (
	"github.com/spf13/cobra"
)

// buildKubeCmd represents the kube command
var buildKubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Creates Kubernetes configuration files.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		return fissile.GenerateKube(flagRoleManifest)

	},
}

func init() {
	buildCmd.AddCommand(buildKubeCmd)
}
