package service

import (
	"fmt"
	"os"

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
