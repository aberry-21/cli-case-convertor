package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode"
)

// sentenceCmd represents the sentence command
var sentenceCmd = &cobra.Command{
	Use:   "sentence",
	Short: "sentence case conversion",
	Args:  cobra.MinimumNArgs(1),
	Long:  `The sentence case converter will allow you to paste any text you’d like, and it will automatically transform it to a fully formed structured sentence. It works by capitalizing the very first letter in each sentence, and will then go on to transform the rest of the text into lowercase as well as converting i’s into I’s. Every letter after a full stop will get converted into an upper case letter.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(convertFunc(args, toSentenceCase))
	},
}

func init() {
	rootCmd.AddCommand(sentenceCmd)
}

func toSentenceCase(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)

	// 1. Capitalize the first letter
	r[0] = unicode.ToUpper(r[0])

	// 2. Convert the rest of the letters to lowercase
	for i := 1; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i]) // to lower case
	}

	return string(r)
}
