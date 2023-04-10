/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"includer/service"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var confiname string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "support by oswaldo@brotherhood",
	Long:  `support by oswaldo@brotherhood`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("package name is null")
		} else {
			service.AddPackage(args[0], confiname)
		}
	},
}

func init() {
	nowpath, _ := os.Getwd()
	patharr := strings.Split(nowpath, "/")
	addCmd.Flags().StringVarP(&confiname, "conf", "c", patharr[len(patharr)-1], "target configure file")
	packageCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
