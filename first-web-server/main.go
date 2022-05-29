package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
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
	fmt.Println("=== 接続を終了します ===")
	conn.Close()
	fmt.Println("=== サーバーを終了します ===")
	// for {
	// 	fmt.Println("=== クライアントからのアクセスを待っています... ===")
	// 	conn, err := listener.Accept()
	// 	fmt.Println("=== クライアントからのアクセスがありました！ ===")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	go func() {
	// 		fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// 		request, err := http.ReadRequest(bufio.NewReader(conn))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		dump, err := httputil.DumpRequest(request, true)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		fmt.Println(string(dump))
	// 		conn.Close()
	// 	}()
	// }
}
