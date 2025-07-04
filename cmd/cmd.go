package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/scottmckendry/pat/ansi"
	"github.com/scottmckendry/pat/img"
)

const VERSION = "0.3.3" // x-release-please-version

var cmd = &cobra.Command{
	Use:   "pat <path-to-image OR url>",
	Short: "like cat, but for pictures",
	Long: `Pat is a tool for displaying images in the terminal.
        It works by converting images to ANSI escape codes that render
        the image in the terminal two pixels at a time (since terminal
        characters are twice as tall as they are wide).`,

	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			cmd.Printf("pat v%s\n", VERSION)
			os.Exit(0)
		}

		if len(args) != 1 {
			// if no arguments are provided, print the usage
			cmd.Usage()
			os.Exit(0)
		}

		path := args[0]
		iamge, err := img.Decode(path)
		cobra.CheckErr(err)

		columns, _ := cmd.Flags().GetInt("columns")
		rows, _ := cmd.Flags().GetInt("rows")

		// assume that if the -r flag is provided but -c is not, the aspect ratio should be preserved
		if rows != 0 && columns == 100 {
			columns = 0
		}

		ansi.PrintImage(iamge, columns, rows)
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
	cmd.Flags().BoolP("version", "v", false, "Display the version of pat and exit")
}
