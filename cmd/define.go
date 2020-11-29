package cmd

import (
	"log"

	"github.com/black-gato/dictionary-cmdl/pkg/define"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(defineCmd)
}

var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "return all information on a word in a languge",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		w := cmd.Flag("word").Value.String()
		if w == "" {
			log.Fatal("word is required")
		}

		l := cmd.Flag("lanaguage").Value.String()
		if l == "" {
			log.Fatal("language is required")
		}

		wrd, err := define.GetWordData(w, l)

	},
}
