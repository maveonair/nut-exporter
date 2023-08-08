package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/maveonair/nut-exporter/internal/nut"
	"github.com/maveonair/nut-exporter/internal/server"
)

func main() {
	portPtr := flag.Int("port", 9055, "Listening Port")
	flag.Parse()

	listeningAddr := fmt.Sprintf("0.0.0.0:%d", *portPtr)

	s := server.NewServer(nut.NewClient())
	log.Fatal(http.ListenAndServe(listeningAddr, s))
}
