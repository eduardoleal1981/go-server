/*
 * Server for static HTML and Files
 */
package main

import (
	"log"
	"net/http"

	"github.com/eduardoleal1981/go-server/server-app/handler"
)

func main() {
	var tHandler handler.THandler
	err := http.ListenAndServe(":80", tHandler)
	// TODO HTTPS ACCESS: log.Fatal(srv.ListenAndServeTLS(certFile, keyFile string))
	if err != nil {
		log.Fatal(err)
	}
}