package main

import (
	"context"
	"fmt"
	"github.com/muhammadkhon-abdulloev/imaginator/config"
	"os"
)

func main() {
	cfg, err := config.Init(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Server.Port)

	files, err := os.ReadDir("assets")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

	}
}
