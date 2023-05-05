package service

import (
	"fmt"
	"includer/tools"
	"io/ioutil"
	"os"
	"os/exec"
)

/*
wait for test
*/
func Build(packname ...string) {
	var err error
	var packid string
	for _, name := range packname {
		packid, err = tools.GetPackId(name)
		if err == nil {
			errlist := build(rootpath + "/lib/" + packid)
			if len(errlist) > 0 {
				PrintFormat(errlist)
			} else {
				fmt.Printf("%s build finished\n", name)
			}
		}
	}
}
func build(dirpath string) (err []error) {
	var newerr error
	fs, newerr := os.Stat(dirpath)
	err = []error{}
	var errlist []error
	if newerr == nil {
		if fs.IsDir() {
			farr, newerr := ioutil.ReadDir(dirpath)
			if newerr == nil {
				for _, fsinfo := range farr {
					errlist = build(dirpath + "/" + fsinfo.Name())
					if len(errlist) > 0 {
						err = append(err, errlist...)
					}
				}
			}
		} else if fs.Name() == "build.sh" {
			//do build
			cmd := exec.Command("bash", "build.sh")
			newerr = cmd.Run()
			if newerr != nil {
				err = append(err, newerr)
			}
		}
	} else {
		err = append(err, newerr)
	}
	return
}

// 格式化输出数组
func PrintFormat[T any](origin []T) {
	if len(origin) > 0 {
		for _, element := range origin {
			fmt.Println(element)
		}
	}
}
