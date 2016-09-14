package main

import (
	"flag"
	"log"
	"net/http"
	"pinyin_server/pinyin"

	rpc "github.com/QianChenglong/rpc/v2"
	"github.com/QianChenglong/rpc/v2/json2"
)

var ListenAddr = "localhost:4000"

func main() {
	flag.Parse()

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json")
	s.RegisterService(new(pinyin.Pinyin), "")

	http.Handle("/", s)
	log.Println("listen at", ListenAddr)
	e := http.ListenAndServe(ListenAddr, nil)
	if e != nil {
		log.Fatal("listen error:", e)
	}
}
