package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/uzixCode/gocode/routes"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		routes.Routes(r)
		port := "5555"
		if len(args) > 0 {
			port = args[0]
		}
		fmt.Printf("http://localhost:%v\n", port)
		r.Run(fmt.Sprintf(":%v", port))
	},
}
