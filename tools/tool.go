package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	include_template = "#include \"repository"
)

func Replace(filepath, old, newdst string) (err error) {
	fe, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err == nil {
		defer fe.Close()
		read := bufio.NewReader(fe)
		buffer := make([]byte, 100<<10)
		lang, err := read.Read(buffer)
		fmt.Println("read finished ,start replace", err)
		if err == nil {
			msg := string(buffer[:lang])
			msg = strings.ReplaceAll(msg, old, newdst)
			fe, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println(err) //debugline
				return err
			}
			_, err = fe.Write([]byte(msg))
			// fmt.Printf("read %v ,content length is %v,after replace length is %v\n", filepath, lang, len(msg))
			fmt.Printf("read %v,write %v\n", lang, len(msg))
			fmt.Println(newdst) //debugline
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
