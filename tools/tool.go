package tools

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/oswaldoooo/octools/toolsbox"
)

const (
	include_template = "#include \"repository"
)

var rootpath = os.Getenv("INCLUDER_HOME")
var linkpath string
var isinit bool = false

func init() {
	if len(rootpath) > 0 {
		linkpath = rootpath + "/conf/link"
		isinit = true
	}
}
func Replace(filepath, old, newdst string) (err error) {
	fe, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err == nil {
		defer fe.Close()
		read := bufio.NewReader(fe)
		buffer := make([]byte, 100<<10)
		lang, err := read.Read(buffer)
		if err == nil {
			msg := string(buffer[:lang])
			msg = strings.ReplaceAll(msg, old, newdst)
			fe, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				return err
			}
			_, err = fe.Write([]byte(msg))
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

// 从获取link表中获取包id
func GetPackId(packname string) (packid string, err error) {
	if !isinit {
		err = errors.New("tools dont init yet")
		return
	}
	resmap, err := toolsbox.ParseList(linkpath)
	if err == nil {
		if packid_, ok := resmap[packname]; ok {
			packid = packid_
		} else {
			err = errors.New(packname + " dont existed in packlink")
		}
	}
	return
}

// 储存包关系进link表中
func SavePackLink(packname, packid string) (err error) {
	_, err = GetPackId(packname)
	if err != nil {
		resmap, err := toolsbox.ParseList(linkpath)
		if err == nil {
			resmap[packname] = packid
			_, err = toolsbox.FormatList(resmap, linkpath)
		}
	} else {
		err = errors.New(packname + " was existed in packlink")
	}
	return
}
