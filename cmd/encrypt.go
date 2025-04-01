package cmd

import (
	"hamster/common"
	"hamster/core"
	"hamster/log"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {

			config := common.GetSystemConfig()
			if config == nil {
				log.Errorf("failed to get config")
			}

			encrypt := core.Encrypt(config.Salt, args[0])

			log.Infof("encrypt data: %s", encrypt)

			decrypt := core.Decrypt(config.Salt, encrypt)

			log.Infof("dencrypt data: %s", decrypt)

		}
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
