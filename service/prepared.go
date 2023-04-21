package service

import (
	"fmt"
	"includer/tools"
	"io/ioutil"
	"os"
	"strings"

	"github.com/oswaldoooo/octools/toolsbox"
)

var rootpath = os.Getenv("INCLUDER_HOME")

// 检查lib和link关系是否对应
func checklink() (err error) {
	resmap, err := toolsbox.ParseList(rootpath + "/conf/link")
	parenpath := rootpath + "/lib/"
	if err == nil {
		for packname, packid := range resmap {
			_, err = os.Stat(parenpath + packid)
			if err != nil {
				//lib不存在，删除lib中的对应关系
				delete(resmap, packname)
				fmt.Printf("package %v dont existed,auto delete it\n", packname)
			}
		}
		_, err = toolsbox.FormatList(resmap, rootpath+"/conf/link")
	}
	return
}

// 替换repository在头文件或cpp文件
func checkheader(dirpath, packid string) {
	fearr, err := ioutil.ReadDir(dirpath)
	if err == nil {
		for _, fe := range fearr {
			if fe.IsDir() && fe.Name()[0] != '.' {
				checkheader(dirpath+"/"+fe.Name(), packid)
			} else {
				if strings.Contains(fe.Name(), ".h") || strings.Contains(fe.Name(), ".cpp") {
					newwords := fmt.Sprintf(include_replace, rootpath+"/lib/"+packid)
					err = tools.Replace(dirpath+"/"+fe.Name(), include_template, newwords)
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}
		}
	}
}
