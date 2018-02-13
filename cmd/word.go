package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(wordCmd)
}

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "convert word into bytes and runes <usage> conv word a",

	Run: func(cmd *cobra.Command, Args []string) {
		arg := Args[0]
		fmt.Println("string - bytes - rune ")
		fmt.Printf("%v\t%v\t%v\t\n", string(arg), []byte(arg), []rune(arg))
	},
}
