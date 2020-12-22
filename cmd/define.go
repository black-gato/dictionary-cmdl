package cmd

import (
	"fmt"
	"log"

	d "dictionary-cmdl/pkg/define"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "return all information on a word in a languge",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("define called")
		l := cmd.Flag("language").Value.String()
		if l == "" {
			log.Fatal("language is required")
		}

		w := cmd.Flag("word").Value.String()
		if w == "" {
			log.Fatal("word is required")
		}

		data, err := d.GetEntry(l, w)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
		// for _, d := range data {
		// 	log.Println(d.Phonetics)

		// }
	},
}

func init() {

	defineCmd.Flags().StringP("language", "l", "", "language you are using")
	defineCmd.Flags().StringP("word", "w", "", "word you are defining")
	if err := viper.BindPFlags(defineCmd.Flags()); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(defineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
