package cmd

import (
	"fmt"
	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

// camelCmd represents the camel command
var camelCmd = &cobra.Command{
	Use:   "camel",
	Short: "camel case conversion",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(convertFunc(args, toCamelCase))
	},
}

func init() {
	rootCmd.AddCommand(camelCmd)
}

func toCamelCase(s string) string {
	return stringy.New(s).CamelCase()
}
