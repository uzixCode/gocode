package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uzixCode/gocode/utils"
)

var scanRoutesCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan Routes",
	Run: func(cmd *cobra.Command, args []string) {
		utils.ScanRoutes("./routes")
	},
}
