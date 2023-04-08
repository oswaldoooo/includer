package service

import (
	"fmt"
	"os"
)

var default_buffer_size = 1 << 10

func Read(fullpath string) {
	fe, err := os.OpenFile(fullpath, os.O_RDONLY, 0600)
	if err == nil {
		defer fe.Close()
		buffer := make([]byte, default_buffer_size)
		lang, err := fe.Read(buffer)
		if err == nil {
			fmt.Println("prepare parser the target")
			err = Parser(buffer[:lang])
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
