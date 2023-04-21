package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"includer/tools"
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
	//预先检查package是否存在
	err = checklink()
	cnfinfo := new(cnf)
	err = xml.Unmarshal(content, cnfinfo)
	if err == nil {
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
	//包名不为空
	if len(include.PackageName) > 0 {
	start:
		linklist, err := toolsbox.ParseList(rootpath + "/conf/link")
		if err == nil {
			if packid, ok := linklist[include.PackageName]; ok {
				//验证包是否存在
				info, err := os.Stat(rootpath + "/lib/" + packid)
				if err != nil || !info.IsDir() {
					//若包不存在或不是一个目录,拉取包
					ok = clonepackage(include.PackageName, packid)
					if ok {
						goto start
					} else {
						fmt.Printf("get package %v failed\n", include.PackageName)
					}
				}

				//搜寻指定的头文件
				headers := []string{}
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
				ok = clonepackage(include.PackageName, "")
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
func clonepackage(packagename, targetname string) bool {
	realpackname := getrepositoryname(packagename)
	err := os.Mkdir(rootpath+"/lib/"+realpackname, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	cmd := exec.Command("git", "clone", "https://"+packagename+".git", rootpath+"/lib/"+realpackname)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Printf("download pacakge %v\n", packagename)
		list, err := toolsbox.ParseList(rootpath + "/conf/link")
		if err == nil {
			respid, err := strconv.Atoi(targetname)
			if len(targetname) == 0 || err != nil {
				//未指定包ID，或之前的包ID被客户串改，则自动生成并记录包ID
				respid = rand.Intn(899999) + 100000
				list[packagename] = strconv.Itoa(respid)
			}
			cmd = exec.Command("mv", rootpath+"/lib/"+realpackname, rootpath+"/lib/"+strconv.Itoa(respid))
			err = cmd.Run()
			if err == nil {
				_, err = toolsbox.FormatList(list, rootpath+"/conf/link")
				linkAllLib(rootpath+"/lib/"+strconv.Itoa(respid), strconv.Itoa(respid))
			}
		}
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
	}
	return true
}

// 从包名中获取最后一位包名
func getrepositoryname(packagename string) (name string) {
	resparr := strings.Split(packagename, "/")
	name = resparr[len(resparr)-1]
	return
}

// 搜寻指定头文件
func searchheader(pathid string, headerlist []string) (err error) {
	resmap := toolsbox.ArrayToMap(headerlist)
	parendir := rootpath + "/lib/" + pathid
	_, err = os.Stat(parendir)
	if err == nil {
		resmap = searchheaderfromdir(resmap, parendir)
		if len(resmap) > 0 {
			errlist := []error{}
			e := 0
			for iname := range resmap {
				errlist = append(errlist, errors.New(iname+" is dont existed"))
				e++
			}
			//collection all errors
			err = errors.Join(errlist...)
		}
	}
	return
}

// 从指定目录中搜寻头文件
func searchheaderfromdir(origin map[string]struct{}, dirpath string) (resmap map[string]struct{}) {
	dirarr, err := ioutil.ReadDir(dirpath)
	if err == nil {
		for _, fileinfo := range dirarr {
			//先进行本目录搜索，若没有再进行深度搜寻
			if fileinfo.IsDir() && !strings.Contains(fileinfo.Name(), "git") {
				//遍历子目录
				origin = searchheaderfromdir(origin, dirpath+"/"+fileinfo.Name())

			} else if _, ok := origin[fileinfo.Name()]; ok {
				//是指定的头文件,建立连接
				nowpath, _ := os.Getwd()
				cmd := exec.Command("ln", dirpath+"/"+fileinfo.Name(), nowpath+"/lib")
				err = cmd.Run()
				if err == nil {
					delete(origin, fileinfo.Name())
				}
			} else {
			}
		}
		resmap = origin
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

// 链接所有动态库至usr_lib目录(待测试),并修改其中的头文件指向仓库地址
func linkAllLib(dirpath string, packid string) {
	fearr, err := ioutil.ReadDir(dirpath)
	if err == nil {
		for _, fe := range fearr {
			if fe.IsDir() {
				linkAllLib(dirpath+"/"+fe.Name(), packid)
			} else {
				if strings.Contains(fe.Name(), ".so") || strings.Contains(fe.Name(), ".dylib") {
					//连接到usr_lib目录，先验证是否存在
					_, err = os.Stat(rootpath + "/usr_lib/" + fe.Name())
					if err != nil {
						cmd := exec.Command("cp", dirpath+"/"+fe.Name(), rootpath+"/usr_lib/"+fe.Name())
						err = cmd.Run()
					} else {
						//已存在不做任何操作
						fmt.Println("library", fe.Name(), "alreay existed")
					}
				} else if strings.Contains(fe.Name(), ".h") || strings.Contains(fe.Name(), ".cpp") { //!!!注，该区域代码并为进行测试，请不要直接运行
					replace_words := fmt.Sprintf(include_replace, rootpath+"/lib/"+packid)
					err = tools.Replace(dirpath+"/"+fe.Name(), include_template, replace_words)
				}
			}
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
