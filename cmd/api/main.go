package main

import (
	"log"
	"os"

	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/config"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	infoLog := log.New(os.Stdout,"INFO\t",log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr,"ERROR\t",log.Ldate | log.Ltime)
	
}
