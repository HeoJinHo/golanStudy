package main

// 크로링 실습
// 대상 사이트 : 루리웹(ruliweb.com)

import (
	_ "bufio"
	"fmt"
	_ "fmt"
	"net/http"
	_ "os"
	_ "strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// 스크래핑 대상 URL
//const urlRoot = "https://www.ruliweb.com"
const urlRoot = "https://www.fmkorea.com/index.php?s_comment_srl=3877955992&mid=football_world&category=853076194"

// 첫번째 방문: 메인페이지 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "title hotdeal_var8"
	}
	return false
}

func errorChecks(e error) {
	if e != nil {
		panic(e)
	}
}

var wg sync.WaitGroup

func main() {

	// 메인 페이지 get 방식 요청
	resp, err := http.Get(urlRoot)
	errorChecks(err)

	// 요청 Body 닫기
	defer resp.Body.Close()

	// 응답 데이터(HTML)
	root, err := html.Parse(resp.Body)
	errorChecks(err)

	// ParseMainNodes 메소드를 크롤링(스크래핑) 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)

	for _, link := range urlList {
		//fmt.Println("index : ", idx, ", Check Link : ", link)
		fmt.Println("TargetURL : https://www.fmkorea.com" + scrape.Attr(link, "href"))
	}

}
