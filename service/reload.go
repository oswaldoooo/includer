package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func Reload(filename string) {
	nowpath, _ := os.Getwd()
	content, err := ioutil.ReadFile(nowpath + "/" + filename)
	if err == nil {
		cf := new(cnf)
		err = xml.Unmarshal(content, cf)
		if err == nil {
			filemap := make(map[string]struct{})
			filemap = appendfiletomap(filemap, nowpath+"/lib")
			for _, include := range cf.Include_Path {
				if len(include.Headers) == 0 {
					//代表链接的是整个目录
					therealname := getrepositoryname(include.PackageName)
					bcinfo, err := os.Stat(nowpath + "/lib/" + therealname)
					if err == nil && bcinfo.IsDir() {
						//代表目录已链接过来，
					} else {
						fmt.Println(include.PackageName, "is not link here")
					}
				} else {
					//单个文件一个一个查找
					comparemap := tomap(include.Headers)
					info, err := os.Stat(nowpath + "/lib")
					if err == nil && info.IsDir() {
						for info := range filemap {
							if _, ok := comparemap[info]; ok {
								//若和includer中的header匹配，则删除其在map中的位置
								delete(comparemap, info)
							}
						}
						if len(comparemap) > 0 {
							toprint := toleaveobj(comparemap)
							fmt.Println(toprint)
						}
					}
				}
			}
		}
	}
}
func toleaveobj(origin map[string]struct{}) (res string) {
	res = ""
	for info := range origin {
		res += info + " is dont exist in lib\n"
	}
	return
}
func tomap(origin []header) (resmap map[string]struct{}) {
	resmap = make(map[string]struct{})
	for _, head := range origin {
		resmap[head.Name] = struct{}{}
	}
	return
}
func appendfiletomap(origin map[string]struct{}, dirpath string) (res map[string]struct{}) {
	fileinfo, err := ioutil.ReadDir(dirpath)
	if err == nil {
		for _, info := range fileinfo {
			if info.IsDir() {
				origin = appendfiletomap(origin, dirpath+"/"+info.Name())
			} else {
				origin[info.Name()] = struct{}{}
			}
		}
	}
	res = origin
	return
}
