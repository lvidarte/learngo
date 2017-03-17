package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

const port string = "5000"
const backendAddress string = "localhost:8080"

var requestBytes map[string]int64
var requestLock sync.Mutex

func init() {
	requestBytes = make(map[string]int64)
}

func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	log.Printf("Proxy listening on port %s", port)
	for {
		if conn, err := ln.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}

// handleConnection is spawned once per connection from a client, and exits when the client is
// done sending requests.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		req, err := http.ReadRequest(reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("Failed to read request: %s", err)
			}
			return
		}

		// Connect to a backend and send the request along.
		if be, err := net.Dial("tcp", backendAddress); err == nil {
			be_reader := bufio.NewReader(be)
			if err := req.Write(be); err == nil {
				if resp, err := http.ReadResponse(be_reader, req); err == nil {
					bytes := updateStats(req, resp)
					resp.Header.Set("X-Bytes", strconv.FormatInt(bytes, 10))

					FixHttp10Response(resp, req)
					if err := resp.Write(conn); err == nil {
						log.Printf("proxied %s: got %d", req.URL.Path, resp.StatusCode)
					}
					if resp.Close {
						return
					}
				}
			}
		}
	}
}

// updateStats takes a request and response and collects some statistics about them. This is
// very simple for now.
func updateStats(req *http.Request, resp *http.Response) int64 {
	requestLock.Lock()
	defer requestLock.Unlock()

	bytes := requestBytes[req.URL.Path] + resp.ContentLength
	requestBytes[req.URL.Path] = bytes
	return bytes
}
