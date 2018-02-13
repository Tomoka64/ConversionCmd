package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(RangeCmd)
}

var RangeCmd = &cobra.Command{
	Use:   "range",
	Short: "convert word into bytes and runes. e.g.) <usage> conv range 50 100",

	Run: func(cmd *cobra.Command, Args []string) {

		fmt.Println("count - string - bytes - rune ")

		i, _ := strconv.Atoi(Args[0])
		j, _ := strconv.Atoi(Args[1])
		for ; i < j; i++ {
			fmt.Printf("%v\t%v\t%v\t%v\t\n", i, string(i), []byte(string(i)), rune(i))

		}
	},
}
