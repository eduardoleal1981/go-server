// Package

package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type THandler int

func (tHandler THandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/js/api") {
		fmt.Println("Debug: tServeAPI")
		tServeAPI(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".js") {
		fmt.Println("Debug: tServeJs")
		tServeJs(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/css/normalize") {
		fmt.Println("Debug: tServeLayout")
		tServeLayout(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".css") {
		fmt.Println("Debug: tServeCss")
		tServeCss(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".svg") {
		fmt.Println("Debug: tServeSvg")
		tServeSvg(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".html") {
		fmt.Println("Debug: tServeHtml .html")
		tServeHtml(w, r)
	} else if strings.HasSuffix(r.URL.Path, "/") {
		fmt.Println("Debug: tServeHtml /")
		tServeHtml(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".ico") {
		fmt.Println("Debug: tServeIco")
		tServeIco(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".jpeg") {
		fmt.Println("Debug: tServeJpeg")
		tServeJpeg(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".png") {
		fmt.Println("Debug: tServePng")
		tServePng(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".pdf") {
		fmt.Println("Debug: tServePdf")
		tServePdf(w, r)
	}
}

func tServeAPI(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	fmt.Println("*** ***")
	fmt.Println("tServeAPI")
	fmt.Println("Opção não usada: public/" + tTrimUrlHost(r.Host) + r.URL.Path)
	fmt.Println("Opção não usada: public/" + r.URL.Path)
	fmt.Println("Opção usada: public" + r.URL.Path)
	http.ServeFile(w, r, "public"+r.URL.Path)
}

func tServeJs(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	// fmt.Println("*** ***")
	fmt.Println("tServeJs")
	fmt.Println("Opção     usada: public/" + tTrimUrlHost(r.Host) + r.URL.Path)
	fmt.Println("Opção não usada: public/" + r.URL.Path)
	fmt.Println("Opção não usada: public" + r.URL.Path)
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+r.URL.Path)
}

func tServeLayout(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	fmt.Println("*** ***")
	fmt.Println("tServeLayout")
	fmt.Println("Opção não usada: public/" + tTrimUrlHost(r.Host) + r.URL.Path)
	fmt.Println("Opção não usada: public/" + r.URL.Path)
	fmt.Println("Opção     usada: public" + r.URL.Path)
	http.ServeFile(w, r, "public"+r.URL.Path)
}

func tServeCss(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	fmt.Println("*** ***")
	fmt.Println("tServeCss")
	fmt.Println("Opção     usada: public/" + tTrimUrlHost(r.Host) + r.URL.Path)
	fmt.Println("Opção não usada: public/" + r.URL.Path)
	fmt.Println("Opção não usada: public" + r.URL.Path)
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+r.URL.Path)
}

func tServeSvg(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve File
	http.ServeFile(w, r, "public/"+r.URL.Path)
}

func tServeHtml(w http.ResponseWriter, r *http.Request) {
	// Protocol
	w.Header().Set("HTTP-Version", "HTTP/2.0") //Chrome OK
	w.Header().Set("X-Firefox-Spdy", "h2")     //Chrome OK
	w.Header().Set("Protocol", "HTTP/2.0")     //Chrome OK
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip") //Chrome OK
	w.Header().Set("Cache-Control", "no-cache") //Chrome OK
	w.Header().Set("Date", time.Now().Format(time.UnixDate))
	// Message
	w.Header().Set("Content-Type", "text/html; charset=utf-8") //Chrome OK
	w.Header().Set("Content-Language", "pt-br")                //Chrome OK
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve File
	fmt.Println("*** ***")
	fmt.Println("tServeHtml")
	fmt.Println("Opção     usada: public/" + tTrimUrlHost(r.Host) + "/index.html")
	fmt.Println("Opção não usada: public/" + "/index.html")
	fmt.Println("Opção não usada: public" + "/index.html")
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/index.html")
}

func tServeIco(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "image/x-icon")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	http.ServeFile(w, r, "public/img/"+r.URL.Path)
}

func tServeJpeg(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "image/jpeg")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+r.URL.Path)
}

func tServePng(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "image/png")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+r.URL.Path)
}

func tServePdf(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "/pdf")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/docs/"+r.URL.Path)
}

func tTrimUrlHost(host string) string {
	var name = strings.TrimSuffix(host, ":8080")
	name = strings.TrimSuffix(name, ".com.br")
	name = strings.TrimSuffix(name, ".com")
	return name
}

// TODO: func database() string
// 	db, err := sql.Open("postgres", "user=postgres dbname=np sslmode=disable")

// TODO: func answer(q Question) Answer
// a, err := ai.Answers(q)
