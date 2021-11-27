package main

import (
	"fmt"
	"log"

	"github.com/smart48ru/Golang-Level1/LESSON8/configuration"
)

func main() {
	cfg, err := configuration.ConfigurationFlag()
	if err != nil {
		fmt.Println(cfg)
		log.Fatalf("error: %v", err)
	}
	fmt.Println(cfg)
}
