package main

import (
	"fmt"
	"github.com/smart48ru/Golang-Level1/LESSON8/configuration"
	"gopkg.in/yaml.v2"
	"log"
)


func main(){
	cfg := configuration.ConfigurationFlag()
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("!! Program configuration !!\n%s\n\n", string(d))

}
