package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "conv",
	Short: "This tool is to convert a string into runes/bytes and vice versa.",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-keisan",
	Long:  `All software has versions. This is go-keisan's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("conv v1.0")
	},
}
