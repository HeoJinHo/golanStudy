package main

import "net/http"

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//data := []byte("Hello World")
	data2 := []byte(test("Test", "1"))
	res.Write(data2)
}

func main() {

	handler := HttpHandler{}
	http.ListenAndServe(":3000", handler)

}

func test(x, y string) string {
	return x + y
}
