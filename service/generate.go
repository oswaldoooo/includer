package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func Generate(filename string) {
	nowpath, _ := os.Getwd()
	newcnf := new(cnf)
	newcnf.XMLName = xml.Name{Space: "", Local: "includer"}
	newcnf.Version = "1.0"
	newcnf.Include_Path = append(newcnf.Include_Path, includers{PackageName: "input the repository name", Headers: []header{header{Name: "header file name"}}})
	by, err := xml.MarshalIndent(newcnf, "", "\t")
	by = append([]byte(xml.Header), by...)
	if err == nil {
		err = ioutil.WriteFile(nowpath+"/"+filename, by, 0600)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
