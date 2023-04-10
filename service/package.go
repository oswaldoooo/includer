package service

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func AddPackage_bak(packagename, path string) {
	nowpath, _ := os.Getwd()
	fe, err := os.OpenFile(nowpath+"/"+path, os.O_RDONLY, 0600)
	if err == nil {
		defer fe.Close()
		buffer := make([]byte, default_buffer_size)
		read := bufio.NewReader(fe)
		lang, err := read.Read(buffer)
		fmt.Println(string(buffer[:lang]))
		// fmt.Print("do c \t", err)
		// fmt.Println("langth is", lang)
		if err == nil || err == io.EOF {
			//解析xml配置文件
			fmt.Print("meet eof ,bug still conitnue") //debugline
			newcnf := new(cnf)
			err = xml.Unmarshal(buffer[:lang], newcnf)
			// fmt.Print(err) //debugline
			if err == nil {
				fmt.Printf("origin data\n%v\n", newcnf)
				// newincluder := includers{PackageName: packagename}
				// include_arr := newcnf.Include_Path
				// include_arr = append(include_arr, newincluder)
				// newcnf.Include_Path = include_arr
				// newcnf.Include_Path = append(newcnf.Include_Path, includers{PackageName: packagename})
				resbytes, err := xml.MarshalIndent(newcnf, "", "\t")
				fmt.Printf("content is \n%v\n", newcnf)
				// fmt.Print(err) //debugline
				if err == nil {
					resbytes = append([]byte(xml.Header), resbytes...)
					ge, err := os.OpenFile(nowpath+"/"+path, os.O_WRONLY|os.O_TRUNC, 0600)
					if err == nil {
						defer ge.Close()
						_, err = ge.Write(resbytes)
						// fmt.Print(err) //debugline
					}
				}
			}
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
func AddPackage(packagename, path string) {
	nowpath, _ := os.Getwd()
	content, err := ioutil.ReadFile(nowpath + "/" + path)
	fmt.Printf("origin data \n%v\n", string(content))
	newcnf := new(cnf)
	err = xml.Unmarshal(content, newcnf)
	// fmt.Print(err) //debugline
	if err == nil {
		// fmt.Printf("origin data\n%v\n", newcnf)
		newcnf.Include_Path = append(newcnf.Include_Path, includers{PackageName: packagename})
		resbytes, err := xml.MarshalIndent(newcnf, "", "\t")
		fmt.Printf("content is \n%v\n", newcnf.Include_Path)
		// fmt.Print(err) //debugline
		if err == nil {
			resbytes = append([]byte(xml.Header), resbytes...)
			ge, err := os.OpenFile(nowpath+"/"+path, os.O_WRONLY|os.O_TRUNC, 0600)
			if err == nil {
				defer ge.Close()
				_, err = ge.Write(resbytes)
				// fmt.Print(err) //debugline
			}
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
