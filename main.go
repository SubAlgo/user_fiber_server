package main

import (
	"github.com/subAlgo/userFiber/database"
	"log"
)

func main() {
	var err error
	if err = database.Connect(); err != nil {
		panic("could not connect to the database")
	}
	if err = database.Migration(); err != nil {
		log.Fatalln(err)
	}
}
