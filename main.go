package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// 1. Initialize our thread-safe store
	store := NewStore()

	// 2. Listen on TCP Port 6379
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Redish (v0.1) listening on port 6379...")

	for {
		// 3. Accept new connections (Blocking call)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		// 4. Handle each connection in a separate Goroutine (Concurrency!)
		go handleConnection(conn, store)
	}
}

func handleConnection(conn net.Conn, store *Store) {
	defer conn.Close()

	// Create a scanner to read the stream line-by-line
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		// Protocol: simple text commands (e.g., "SET key value")
		commandLine := scanner.Text()
		parts := strings.Fields(commandLine)

		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToUpper(parts[0])

		switch cmd {
		case "SET":
			if len(parts) < 3 {
				conn.Write([]byte("ERR usage: SET key value\n"))
				continue
			}
			key := parts[1]
			val := parts[2]
			store.Set(key, val)
			conn.Write([]byte("OK\n"))

		case "GET":
			if len(parts) < 2 {
				conn.Write([]byte("ERR usage: GET key\n"))
				continue
			}
			key := parts[1]
			val, found := store.Get(key)
			if !found {
				conn.Write([]byte("(nil)\n"))
			} else {
				conn.Write([]byte(val + "\n"))
			}

		default:
			conn.Write([]byte("ERR unknown command\n"))
		}
	}
}
