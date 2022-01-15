package main

import (
	"log"
	"task1/internal/server"
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