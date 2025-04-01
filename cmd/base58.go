package cmd

import (
	"hamster/common"
	"hamster/core"
	"hamster/log"
	"strings"

	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
)

var base58Cmd = &cobra.Command{
	Use:   "base58",
	Short: "base58",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {

			config := common.GetSystemConfig()
			if config == nil {
				log.Errorf("failed to get config")
			}

			if !strings.Contains(args[0], "\"") {
				log.Errorf("string must be a json bytes")
				return
			}

			secret := []byte{}
			err := core.FastJsonUnmarshalFromString(args[0], &secret)
			if err != nil {
				log.Errorf("can not unmarshal json bytes: %s", err.Error())
				return
			}

			privateKey := base58.Encode(secret)
			log.Infof("privateKey: %s", privateKey)

		}
	},
}

func init() {
	rootCmd.AddCommand(base58Cmd)
}
