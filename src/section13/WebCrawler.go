// 최종 실습 예제
// 대상 사이트 : 클리앙 소모임 (https://www.clien.net/service/somoim)

package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"
	"os"
	_ "os"
	"strings"
	_ "strings"
	"sync"
	_ "sync"

	_ "github.com/yhat/scrape" // Scraping Package

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	_ "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	_ "golang.org/x/net/html/atom" // DOM Mapping Package
)

// Package 설치 - go get

// Scraping Target URL
const (
	urlRoot       = "https://www.clien.net"
	urlSomoimList = "https://www.clien.net/service/somoim"
)

// Check Error in Common
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// 첫 번째 방문 대상으로 원하는 URL을 파싱 후 반환
func parseURLList(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n, "class") == "somoim_item"
	}

	return false
}

// 소모임명 을 파싱 후 반환
func parseSomainName(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n, "class") == "subject_fixed"
	}

	return false
}

// URL 대상이 되는 페이지(서브 페이지) 대상으로 원하는 내용을 파싱 후 반환
func scrapContents(fileName, url string) {
	defer wg.Done() // 작업 종료 알림

	// Get 방식 요청
	resp, err := http.Get(url)
	checkError(err)

	// 요청 Body 닫기
	defer resp.Body.Close()

	// 응답 데이터(HTML)
	root, err := html.Parse(resp.Body)
	checkError(err)

	// Response 데이터(HTML) 의 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.Span && scrape.Attr(n, "class") == "subject_fixed"
	}

	// 파일 스트림 생성(열기) -> 파일명, 옵션 권한
	file, err := os.OpenFile("./scrape/"+fileName, os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	checkError(err)

	// 메소드 종료 시 파일 닫기
	defer file.Close()

	// 쓰기 버퍼 선언
	w := bufio.NewWriter(file)

	// matchNode 메소드를 사용해서 원하는 노드 순회(Iterator) 하면서 출력
	for _, g := range scrape.FindAll(root, matchNode) {
		// url 및 해당 데이터 출력
		// fmt.Println("result : ", scrape.Text(g))

		// 파싱 데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}
	w.Flush()
}

// 동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

func main() {
	// Step 1. 소모임에 대한 Main Page GET 방식 요청
	resp, err := http.Get(urlSomoimList)
	checkError(err)

	// Release Resource
	defer resp.Body.Close()

	// response.data (html)
	root, err := html.Parse(resp.Body)
	checkError(err)

	// ParseBoardList : 각 게시판 URL 추출
	urlList := scrape.FindAll(root, parseURLList)

	// 대상 URL 추출
	for idx, link := range urlList {
		somoimTitle := strings.Split(scrape.Text(link), " ")[0]
		somoimUrl := scrape.Attr(link, "href")

		fileName := fmt.Sprintf("%03d_%s.txt", idx+1, somoimTitle)
		// fmt.Printf("%s - %s\n", fileName, somoimUrl)
		// fmt.Printf("%03d - %s - %s\n", idx, SomoimTitle, somoinUrl)

		// 작업 대기열에 추가
		wg.Add(1) // Done 개수와 일치

		// goroutine 시작 -> 작업 대기열 개수와 같아야 한다.
		go scrapContents(fileName, urlRoot+somoimUrl)
	}

	// 모든 작업이 종료될 떄 까지 대기
	wg.Wait()

	// Step 2. Main Page 에서 각 게시판 URL 가져오기

}
