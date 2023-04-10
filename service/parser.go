package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/oswaldoooo/octools/toolsbox"
)

type cnf struct {
	XMLName xml.Name `xml:"includer"`
	// Description  string      `xml:",innerxml"`
	Version      string      `xml:"version,attr"`
	Include_Path []includers `xml:"include_config"`
}
type includers struct {
	XMLName     xml.Name `xml:"include_config"`
	PackageName string   `xml:"packagename,attr"`
	Headers     []header `xml:"header"`
}
type header struct {
	XMLName xml.Name `xml:"header"`
	Name    string   `xml:"name,attr"`
}

func Parser(content []byte) (err error) {
	cnfinfo := new(cnf)
	// fmt.Println("goto parser body")
	err = xml.Unmarshal(content, cnfinfo)
	if err == nil {
		// fmt.Println("do parser")
		for _, include := range cnfinfo.Include_Path {
			includeparser(&include)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
func includeparser(include *includers) {
	// fmt.Printf("start into includer parser %v\n", include.PackageName) //debugline
	//包名不为空
	if len(include.PackageName) > 0 {
	start:
		linklist, err := toolsbox.ParseList(rootpath + "/conf/link")
		if err == nil {
			if _, ok := linklist[include.PackageName]; ok {
				//搜寻指定的头文件
				headers := []string{}
				// fmt.Printf("the header length is %v\n", len(include.Headers)) //debugline
				if len(include.Headers) > 0 {
					for _, heads := range include.Headers {
						headers = append(headers, heads.Name)
					}
					fmt.Println("start search header file") //debugline
					err = searchheader(linklist[include.PackageName], headers)
				} else {
					//默认将整个目录链接到当前lib目录下
					nowpath, _ := os.Getwd()
					cmd := exec.Command("ln", "-s", rootpath+"/lib/"+linklist[include.PackageName], nowpath+"/lib/"+getrepositoryname(include.PackageName))
					err = cmd.Run()
				}
			} else {
				//不存在，则拉取包
				ok = clonepackage(include.PackageName)
				if ok {
					goto start
				} else {
					fmt.Printf("get package %v failed\n", include.PackageName)
				}
			}
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Println("package name is null")
	}
}
func clonepackage(packagename string) bool {
	realpackname := getrepositoryname(packagename)
	err := os.Mkdir(rootpath+"/lib/"+realpackname, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	cmd := exec.Command("git", "clone", "https://"+packagename+".git", rootpath+"/lib/"+realpackname)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	// fmt.Println("start clone", packagename) //debugline
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Printf("download pacakge %v\n", packagename)
		list, err := toolsbox.ParseList(rootpath + "/conf/link")
		if err == nil {
			respid := rand.Intn(899999) + 100000
			list[packagename] = strconv.Itoa(respid)
			cmd = exec.Command("mv", rootpath+"/lib/"+realpackname, rootpath+"/lib/"+strconv.Itoa(respid))
			err = cmd.Run()
			if err == nil {
				_, err = toolsbox.FormatList(list, rootpath+"/conf/link")
			}
		}
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
	}
	return true
}
func getrepositoryname(packagename string) (name string) {
	resparr := strings.Split(packagename, "/")
	name = resparr[len(resparr)-1]
	return
}

// 搜寻指定头文件
func searchheader(pathid string, headerlist []string) (err error) {
	resmap := toolsbox.ArrayToMap(headerlist)
	//debugzone
	// fmt.Printf("list is %v\nmap is %v\n", headerlist, resmap)

	//
	parendir := rootpath + "/lib/" + pathid
	_, err = os.Stat(parendir)
	if err == nil {
		// fmt.Printf("the result map's length is %v,resmap is %v\n", len(resmap), resmap) //debugline
		resmap = searchheaderfromdir(resmap, parendir)
		// fmt.Printf("the result map's length is %v,resmpa is %v\n", len(resmap), resmap)
		if len(resmap) > 0 {
			errlist := []error{}
			e := 0
			// fmt.Println(resmap) //debugline
			for iname := range resmap {
				errlist = append(errlist, errors.New(iname+" is dont existed"))
				e++
			}
			// fmt.Println(len(resmap)) //debugline
			//collection all errors
			err = errors.Join(errlist...)
		}
	}
	return
}

// 从指定目录中搜寻头文件
func searchheaderfromdir(origin map[string]struct{}, dirpath string) (resmap map[string]struct{}) {
	// fmt.Printf("origin is %v\n", origin)
	dirarr, err := ioutil.ReadDir(dirpath)
	if err == nil {
		// fmt.Println("do here") //debugline
		for _, fileinfo := range dirarr {
			//先进行本目录搜索，若没有再进行深度搜寻
			if fileinfo.IsDir() && !strings.Contains(fileinfo.Name(), "git") {
				//遍历子目录
				// fmt.Printf("did here,next directory %v\n", fileinfo.Name())
				origin = searchheaderfromdir(origin, dirpath+"/"+fileinfo.Name())

			} else if _, ok := origin[fileinfo.Name()]; ok {
				//是指定的头文件,建立连接
				nowpath, _ := os.Getwd()
				// fmt.Println("prepare to do command \n", "ln", dirpath+"/"+fileinfo.Name(), nowpath+"/lib") //debugline
				cmd := exec.Command("ln", dirpath+"/"+fileinfo.Name(), nowpath+"/lib")
				err = cmd.Run()
				if err == nil {
					// fmt.Println("did delete act")
					delete(origin, fileinfo.Name())
				}
			} else {
				// fmt.Println("dont match below") //debugline
				// fmt.Println(resmap)
			}
		}
		resmap = origin
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
