package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var (
	SERVER   bool
	ADDR     string
	TRIES    = 3
	upgrader = websocket.Upgrader{}
)

func main() {
	flag.BoolVar(&SERVER, "s", false, "host or connect to remote")
	flag.StringVar(&ADDR, "addr", ":8080", "remote host address")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	if SERVER {
		mux := http.NewServeMux()
		mux.HandleFunc("/", Server)
		fmt.Printf("Listening on %s\n", ADDR)
		log.Fatal(http.ListenAndServe(ADDR, mux))
	} else {
		u := url.URL{Scheme: "ws", Host: ADDR, Path: "/"}
		fmt.Printf("Connecting to %s\n", ADDR)
		Client(u)
	}
}

func Server(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("upgrade: %v", err)
	}
	defer conn.Close()

	Game(conn)
}

func Client(u url.URL) {
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	Game(conn)
}
