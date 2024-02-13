package cmd

import (
	"fmt"
	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

// snakeCmd represents the snake command
var snakeCmd = &cobra.Command{
	Use:   "snake",
	Short: "snake case conversion",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(convertFunc(args, toSnakeCase))
	},
}

func init() {
	rootCmd.AddCommand(snakeCmd)
}

func toSnakeCase(s string) string {
	return stringy.New(s).SnakeCase().ToLower()
}
