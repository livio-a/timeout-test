package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := 8080
	idleTimeout := 1 * time.Millisecond
	readTimeout := 1 * time.Millisecond
	writeTimeout := 1 * time.Millisecond

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %v", r.URL.Path)
	})

	http2Server := &http2.Server{IdleTimeout: idleTimeout}
	http1Server := &http.Server{Handler: h2c.NewHandler(handler, http2Server), ReadTimeout: readTimeout, IdleTimeout: idleTimeout, WriteTimeout: writeTimeout}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	err = http1Server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
