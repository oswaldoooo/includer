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

var rootpath = os.Getenv("INCLUDER_HOME")

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init project",
	Long:  `init includer file by target config file`,
	Run: func(cmd *cobra.Command, args []string) {
		nowpath, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//检查lib目录是否存在，不存在则创建
		info, err := os.Stat(nowpath + "/lib")
		if err != nil {
			os.Mkdir(nowpath+"/lib", 0755)
		} else if err == nil && !info.IsDir() {
			fmt.Println("lib is not directory")
			return
		}
		// fmt.Println("do prepare read") //debugline
		if len(args) == 0 {
			//默认执行和父目录同名的配置文件
			patharr := strings.Split(nowpath, "/")
			default_name := patharr[len(patharr)-1] + ".xml"
			service.Read(nowpath + "/" + default_name)
		} else {
			for _, conf := range args {
				if !strings.Contains(conf, ".xml") {
					//自动添加.xml厚嘴
					conf += ".xml"
				}
				// fmt.Println("prepare read ", nowpath+"/"+conf) //debugline
				service.Read(nowpath + "/" + conf)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
