package main

import (
	"clean-go/common"
	"clean-go/presentation/http"
)

func main() {
	config, err := common.GetInstance()

	if err != nil {
		panic(err)
	}

	container, err := common.NewContainer(config)

	if err != nil {
		panic(err)
	}

	server := http.NewServer(container)
	server.Listen("0.0.0.0", 3000)
}
