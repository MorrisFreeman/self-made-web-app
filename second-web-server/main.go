package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== サーバーを起動します ===")
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("=== クライアントからのアクセスを待っています... ===")
	conn, err := listener.Accept()
	fmt.Println("=== クライアントからのアクセスがありました！ ===")
	// このタイミングではtcpのコネクションが接続されているだけ
	// httpリクエストの読み込みはまだ
	if err != nil {
		panic(err)
	}
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	fmt.Println("=== レスポンスを返します ===")
	header := http.Header{}
	header.Add("Content-Type", "text/html; charset=utf-8")
	header.Add("Set-Cookie", "foo=bar")
	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     header,
		Body:       io.NopCloser(strings.NewReader("Hello World\n")),
	}
	response.Write(conn)
	response.Write(os.Stdout)
	fmt.Println("=== 接続を終了します ===")
	conn.Close()
	fmt.Println("=== サーバーを終了します ===")
}
