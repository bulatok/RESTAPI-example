package main

import (
	"github.com/bulatok/task/internal/server"
	"log"
)

func main(){
	conf, err := server.NewConfig("config.yml")
	if err != nil{
		log.Fatal(err)
	}
	if err := server.Start(conf); err != nil{
		log.Fatal(err)
	}
}