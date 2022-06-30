// Panic & Recover(4)

package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	// defer 함수 (Panic 호출 시 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error: ", r)
		}
	}()

	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("fileOpen : ", f.Name())
	}

	defer f.Close()
}

func main() {
	// Panic & Recover - 최종 실습

	fileOpen("unknown")
}
