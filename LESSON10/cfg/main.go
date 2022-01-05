package main

import (
	"fmt"
	"log"

	"github.com/smart48ru/Golang-Level1/LESSON10/cfg/configuration"
)

func main() {
	cfg, err := configuration.GetConfiguration()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	cfgString,err  := cfg.String()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(cfgString)
}
