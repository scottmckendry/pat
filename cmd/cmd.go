package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"pat/ansi"
	"pat/img"
)

var cmd = &cobra.Command{
	Use:   "pat <path-to-image>",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		if !PathExists(path) {
			cobra.CheckErr(fmt.Errorf("could not find file: %s", path))
		}

		image, err := img.Decode(path)
		cobra.CheckErr(err)

		columns, _ := cmd.Flags().GetInt("columns")
		rows, _ := cmd.Flags().GetInt("rows")

		// assume that if the -r flag is provided but -c is not, the aspect ratio should be preserved
		if rows != 0 && columns == 100 {
			columns = 0
		}

		ansi.PrintImage(image, columns, rows)
	},
}

func Execute() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cmd.Flags().IntP("columns", "c", 100, "Number of columns to use for the image")
	cmd.Flags().IntP("rows", "r", 0, "Number of rows to use for the image")
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
