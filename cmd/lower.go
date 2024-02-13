package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// lowerCmd represents the lower command
var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "lower case conversion",
	Args:  cobra.MinimumNArgs(1),
	Long:  `If you are wondering how to uncapitalize text, this is exactly what the lower case text converter will allow you to do - it transforms all the letters in your text into lowercase letters.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(convertFunc(args, strings.ToLower))
	},
}

func init() {
	rootCmd.AddCommand(lowerCmd)
}
