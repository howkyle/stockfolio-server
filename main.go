package main

import "github.com/howkyle/stockfolio-server/server"

func main() {
	server := server.Create(":8000")
	server.Start()
}
