package main

import (
	"clean-go/common"
	"clean-go/presentation/http"
	"fmt"
)

func main() {
	container, err := common.NewContainer()

	if err != nil {
		fmt.Println(err)
	} else {
		server := http.NewServer(container)

		server.Listen("0.0.0.0", 3000)
	}
}
