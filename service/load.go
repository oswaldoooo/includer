package service

import (
	"errors"
	"fmt"
	"os"

	"github.com/oswaldoooo/octools/toolsbox"
)

// 通过名称加载
func LoadByName(packname string) {
	resmap, err := toolsbox.ParseList(rootpath + "/conf/link")
	if err == nil {
		if packid, ok := resmap[packname]; ok {
			LoadById(packid)
		} else {
			err = errors.New("package " + packname + " dont existed")
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 通过id加载
func LoadById(packid string) {
	parentpath := rootpath + "/lib/" + packid
	_, err := os.Stat(parentpath)
	if err == nil {
		checkheader(parentpath, packid)
	} else {
		fmt.Println(err.Error())
	}

}
