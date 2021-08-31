package main

// 크로링 실습
// 대상 사이트 : 루리웹(ruliweb.com)

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
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

// URL 대상이 되는 페이지 대상으로 원하는 내용을 파싱 후 반환
func scrapeContents(url, fn string) {
	defer wg.Done()

	resp, err := http.Get(url + fn)
	errorChecks(err)

	defer resp.Body.Close()

	// 응답 데이터
	root, err := html.Parse(resp.Body)
	errorChecks(err)

	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.Span && scrape.Attr(n, "class") == "np_18px_span"
	}
	fileName := "/Users/jhheo/Desktop/scrape/" + fn + ".txt"
	// 파일명 생성
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errorChecks(err)

	defer file.Close()

	// bufer
	w := bufio.NewWriter(file)

	// mathnode method 사용 해서 원하는 노드 순회(이터레이터)
	for _, g := range scrape.FindAll(root, matchNode) {
		// url 및 해당 데이터 출력
		fmt.Println("Resault : ", url+fn, scrape.Text(g))
		// 파싱 데이터 -> 버퍼에 기록
		//ioutil.WriteFile(fileName, []byte(scrape.Text(g) + "\r\n"), os.FileMode(0777))
	}

	w.Flush()

}

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

	for i, _ := range urlList {
		//fmt.Println("index : ", idx, ", Check Link : ", link)
		//fmt.Println("TargetURL : https://www.fmkorea.com" + scrape.Attr(link, "href"))
		//fileName := strings.Replace(scrape.Attr(link, "href"), "", "", 1)
		//fileName := scrape.Attr(link, "href")
		wg.Add(1)

		go scrapeContents("https://www.fmkorea.com", string(i))
	}
	// 모든 작업이 끝날때까지 대기
	wg.Wait()

}
