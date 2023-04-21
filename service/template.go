package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	build_template = "#!/bin/bash\nread -p \"input build filename(dont include .cpp)>>\" filename\ng++ -std=c++17 -o $filename $filename.cpp -L$INCLUDER_HOME/usr_lib -g"
	//need input replace path
	replace_template = "\"s/#include \\\"repository/#include \\\"%v/g\""
	include_template = "#include \"repository"
	include_replace  = "#include \"%v"
)

func GenerateStart() {
	nowpath, err := os.Getwd()
	if err == nil {
		_, err = os.Stat(nowpath + "/start.sh")
		if err != nil {
			err = ioutil.WriteFile(nowpath+"/start.sh", []byte(build_template), 0644)
		} else {
			err = errors.New("start.sh already existed")
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
