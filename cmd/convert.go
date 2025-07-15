/*
Copyright © 2025 Manish Ramanandi <bairagimanish222@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/manishramanandi/md2pdf-go-cli/util"

)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
    Use:   "convert",
    Short: "The commands generates the PDF from a given Markdown file",
    Long: `The command accepts a Markdown file and a corresponding PDF file name. It converts the file to PDF.
    For example:

The cli will use the md2pdf library to convert your given Markdown and convert it to a PDF with
your given arguments.`,
    Run: func(cmd *cobra.Command, args []string) {
        mdFile := viper.GetString("mdfile")
        pdfFile := viper.GetString("pdffile")

        // Call util function to convert the files
        err := util.ConvertFile(mdFile, pdfFile)
        if err != nil {
            fmt.Println("Error converting file:", err)
            os.Exit(1)
        }

        fmt.Println("Converted %s to %s ... \n", mdFile, pdfFile)
    },
}


func init() {
	//Define falgs for convert command
	convertCmd.Flags().StringP("mdfile", "m", "", "Markdown file name")
	convertCmd.Flags().StringP("pdffile", "p", "", "PDF file name")


	convertCmd.Flags().StringP("fontsize", "f", "12", "The font size of the pdf")
	convertCmd.Flags().StringP("fontstyle", "s", "B", "The font style of the pdf")
	convertCmd.Flags().StringP("textcolorR", "r", "200", "The text color of the PDF")
	convertCmd.Flags().StringP("textcolorG", "g", "200", "The text color of the PDF")
	convertCmd.Flags().StringP("textcolorB", "b", "200", "The text color of the PDF")


	// Bind flags to viper configuration
	viper.BindPFlag("mdfile", convertCmd.Flags().Lookup("mdfile"))
	viper.BindPFlag("pdffile", convertCmd.Flags().Lookup("pdffile"))

	viper.BindPFlag("fontsize", convertCmd.Flags().Lookup("fontsize"))
	viper.BindPFlag("fontstyle", convertCmd.Flags().Lookup("fontstyle"))
	viper.BindPFlag("textcolorR", convertCmd.Flags().Lookup("textcolorR"))
	viper.BindPFlag("textcolorG", convertCmd.Flags().Lookup("textcolorG"))
	viper.BindPFlag("textcolorB", convertCmd.Flags().Lookup("textcolorB"))

	// Add convert command as a subcommand
	rootCmd.AddCommand(convertCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
