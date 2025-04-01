package main

import (
	"fmt"
	"hamster/cmd"

	"github.com/spf13/cobra"
)

var version = ""
var buildDate = ""
var gitCommit = ""
var gitCommitTime = ""
var goVersion = ""

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version:", version)
		fmt.Println("buildDate:", buildDate)
		fmt.Println("gitCommit:", gitCommit)
		fmt.Println("gitCommitTime:", gitCommitTime)
		fmt.Println("goVersion:", goVersion)
	},
}

func init() {
	cmd.AddCommand(versionCmd)
}

func main() {
	cmd.Execute()
}
