package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent", "hogehoge")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Set("Sec-Fetch-Site", "none")
	request.Header.Set("Sec-Fetch-Mode", "navigate")
	request.Header.Set("Sec-Fetch-User", "?1")
	request.Header.Set("Sec-Fetch-Dest", "document")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept-Language", "ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7")
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
