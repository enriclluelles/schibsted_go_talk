package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/enriclluelles/schibsted_go_talk/wsserver"
)

var addr = flag.String("addr", ":8080", "http service address")
var cert = flag.String("cert", "./misc/cert.pem", "certificate")
var key = flag.String("key", "./misc/key.pem", "key")

func main() {
	flag.Parse()
	s := wsserver.NewServer()
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/ws", s)
	err := http.ListenAndServeTLS(*addr, *cert, *key, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
