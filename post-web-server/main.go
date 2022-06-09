package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== サーバーを起動します ===")
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
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
			response, err := buildResponse(request)
			if err != nil {
				panic(err)
			}
			response.Write(conn)

			fmt.Println("=== 接続を終了します ===")
			conn.Close()
		}()
	}
}

func buildResponse(request *http.Request) (http.Response, error) {
	file, err := func() (io.Reader, error) {
		switch request.URL.Path {
		case "/", "index", "index.html":
			return os.Open("public/index.html")
		case "/now":
			return strings.NewReader(time.Now().String()), nil
		case "/show_request":
			return buildShowRequestHtml(request)
		case "/form", "/form.html":
			return os.Open("public/form.html")
		default:
			return os.Open("public" + request.RequestURI)
		}
	}()
	if err != nil {
		return responseNotFound(), nil
	}
	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     buildHeader(request, strconv.Itoa(100)),
		Body:       io.NopCloser(file),
	}
	return response, nil
}

func buildHeader(request *http.Request, size string) http.Header {
	header := http.Header{}
	if strings.Contains(request.RequestURI, "png") {
		header.Add("Content-Type", "image/png")
	} else if strings.Contains(request.RequestURI, "css") {
		header.Add("Content-Type", "text/css; charset=utf-8")
	} else {
		header.Add("Content-Type", "text/html; charset=utf-8")
	}
	header.Add("Date", time.Now().String())
	header.Add("Content-Length", size)
	return header
}

func responseNotFound() http.Response {
	return http.Response{
		StatusCode: 404,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       io.NopCloser(strings.NewReader("404")),
	}
}

func buildShowRequestHtml(request *http.Request) (io.Reader, error) {
	reqestLine := fmt.Sprintf("%s %s %s", request.Method, request.RequestURI, request.Proto)
	headerHtml := ""
	for k, v := range request.Header {
		headerHtml += k + ": " + strings.Join(v, "; ")
		headerHtml += "<br>"
	}
	// headerByte, _ := json.Marshal(request.Header)
	// header := string(headerByte)
	body, err := io.ReadAll(request.Body)
	html := `<html>
	<header>
	</header>
	<body>
		<h1>Show Request</h1>
		<h2>Request Line</h2>
		<div>
	` + reqestLine + `
		</div>
		<h2>Header</h2>
		<div>
	` + headerHtml + `
		</div>
		<h2>Body</h2>
		<div>
	` + string(body) + `
		</body>
		</html>
	`
	return strings.NewReader(html), err
}
