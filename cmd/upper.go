package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "upper case conversion",
	Args:  cobra.MinimumNArgs(1),
	Long:  `The upper case transformer will take any text that you have and will generate all the letters into upper case ones. It will essentially make all lower case letters into CAPITALS (as well as keep upper case letters as upper case letters).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(convertFunc(args, strings.ToUpper))
	},
}

func init() {
	rootCmd.AddCommand(upperCmd)
}
