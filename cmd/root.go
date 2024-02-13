package cmd

import (
	"github.com/spf13/cobra"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "cli-case-convertor",
	Short: "A cli tool for case conversion",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func convertFunc(args []string, convertor func(s string) string) string {
	ret := strings.Builder{}
	for _, arg := range args {
		ret.WriteString(convertor(arg))
		ret.WriteRune(' ')
	}
	return ret.String()
}
