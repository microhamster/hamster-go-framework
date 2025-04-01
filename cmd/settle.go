package cmd

import (
	"hamster/core"
	settleServer "hamster/servers/settle"

	"github.com/spf13/cobra"

	"hamster/log"
)

var settleServerCmd = &cobra.Command{
	Use:   "settle",
	Short: "Start settle handler",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("event:%s msg:service_exception %v caller_stack:%s",
					core.EVENT_APP_PANIC, r, core.GetCallerStackLog())
			}
		}()

		server := settleServer.NewSettleServer()
		core.Run(server)

	},
}

func init() {
	rootCmd.AddCommand(settleServerCmd)
}
