/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"includer/service"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load header or library",
	Long:  `load header or library`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			//load all
			fearr, err := ioutil.ReadDir(rootpath + "/lib")
			if err == nil {
				for _, fe := range fearr {
					if fe.IsDir() {
						service.LoadById(fe.Name())
					} else {
						fmt.Printf("%v is not directory\n", fe.Name())
					}
				}
			}
		} else {
			for _, name := range args {
				service.LoadByName(name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
