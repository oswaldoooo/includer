/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"includer/service"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Long:  `generate the config file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			nowpath, _ := os.Getwd()
			patharr := strings.Split(nowpath, "/")
			service.Generate(patharr[len(patharr)-1] + ".xml")
		} else {
			for _, name := range args {
				if strings.Contains(name, ".xml") {
					service.Generate(name)
					service.GenerateStart()
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
