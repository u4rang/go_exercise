// File IO (2)
package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
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
	// File IO - ioutil 패키지 및 bufio 활용
	// - 파일 읽기, 버퍼 사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, Writer, Read 메소드를 사용 가능
	/*
		    type Reader interface {
				Read(p []byte) (n int, err error)
			}
			type writer interface {
				write(p []byte) (n int, err error)
			}
	*/
	// 즉, bufio의 NewReader, NewWriter를 통해 객체 반환 후 메서드 사용

	// bufio(Buffered io) 패키지
	// https://golang.org/pkg/bufio

	// 파일 열기
	// 두 번째 매개변수 확인
	// https://golang.org/pkg/os/#pkg-contants

	/*
		상태
		a ---- > a
		b ---- > ab
		c ---- > abc
		d ---- > abcd
		e ---- > e -----> abcd
		f ---- > ef -----> abcd
		g ---- > efg -----> abcd
		h ---- > efgh -----> abcd
		i ---- > i -----> abcdefgh
	*/

	file, err := os.OpenFile("../test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	// bufio 파일 쓰기 예제
	// wt := bufio.NewWriter(w)
	wt := bufio.NewWriter(file) // Writer 변환(버퍼 사용)
	wt.Write([]byte("Hello, world!\nFile Write Test! - 1\n"))
	wt.Write([]byte("Hello, world!\nFile Write Test! - 2\n"))

	// Check Error
	checkErr(err)

	// 버퍼 정보 출력
	fmt.Printf("사용한 Buffer Size : %d bytes\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size : %d bytes\n", wt.Available())
	fmt.Printf("전체 Buffer Size : %d bytes\n", wt.Size())

	wt.Flush() // 버퍼 비우고 반영(버퍼에 내용을 디스크에 기록)

	fmt.Println("쓰기 작업 완료")
	fmt.Println("===============================")

	rt := bufio.NewReader(file) // Reader 반환
	fl, err := file.Stat()
	checkErr(err)

	b := make([]byte, fl.Size())

	fmt.Println("파일 정보 출력(1) : ", fl)
	fmt.Println("파일 이름(2) : ", fl.Name())
	fmt.Println("파일 크기(3) : ", fl.Size())
	fmt.Println("파일 수정 시간(4) : ", fl.ModTime())
	fmt.Println("==================")

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b) // 읽기(ReadLine, ReadByte, ReadBytes 등)

	fmt.Printf("전체 Buffer Size : %d bytes\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : %d bytes\n", data)
	fmt.Println("=================")
	fmt.Println(string(b))
	fmt.Println("=================")

	defer file.Close()

}
