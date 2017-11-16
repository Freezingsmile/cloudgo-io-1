package main

import (
	"cloudgo-io/server"
)

func main() {
	s := server.NewServer()
	s.Run(":8080")
}
