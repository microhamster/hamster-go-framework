package cmd

import (
	"fmt"
	"hamster/core"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hamster",
	Short: "hamster",
	Long:  `hamster`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func AddCommand(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
}

func init() {
	cobra.OnInitialize(core.InitConfig)
	rootCmd.PersistentFlags().StringVarP(&core.CfgFile, "config", "c", "", "config file (./main.yaml)")
}
