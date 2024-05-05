package main

import (
	"clean-go/common"
	"clean-go/presentation/http"
)

func main() {
	container := common.NewContainer()
	server := http.NewServer(container)

	server.Listen("0.0.0.0", 3000)
}
