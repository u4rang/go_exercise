// File IO (1)
package main

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	_ "io/ioutil"
	"os"
	_ "os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// File IO - ioutil 패키지 활용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	// 다음 메소드 확인 가능
	// WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/pkg/io/ioutil/

	s := "Hello, golang!\n File Write Test!\n"

	// File Mode (chmod, chown, chgrp) - Permission
	// 읽기 : 4, 쓰기 : 2, 실행 : 1
	// 소유자, 그룹, 기타 사용자 순서 (777)
	// https://golang.org/pkg/os/#FileMode

	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	checkErr(err)

	data, err := ioutil.ReadFile("../sample.txt")
	checkErr(err)

	fmt.Println("----------------------------")
	fmt.Println(string(data))
	fmt.Println("----------------------------")

}
