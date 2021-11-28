package main

import (
	"fmt"
	"log"

	"github.com/smart48ru/Golang-Level1/LESSON9/configuration"
)

func main() {
	cfg, err := configuration.GetConfiguration()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(cfg.String())
}
