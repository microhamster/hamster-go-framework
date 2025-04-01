package cmd

import (
	"hamster/core"
	apiServer "hamster/servers/api"

	"github.com/spf13/cobra"

	"hamster/log"
)

var apiServerCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api handler",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("event:%s msg:service_exception %v caller_stack:%s",
					core.EVENT_APP_PANIC, r, core.GetCallerStackLog())
			}
		}()

		server := apiServer.NewApiServer()
		core.Run(server)

	},
}

func init() {
	rootCmd.AddCommand(apiServerCmd)
}
