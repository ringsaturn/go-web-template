package main

import (
	"errors"
	"log"
	"os"

	"github.com/ringsaturn/go-web-template/pkg/config"
)

func main() {
	prjPath := os.Getenv("PRJ_PATH")
	if prjPath == "" {
		panic(errors.New("PRJ_PATH required, try\nexport PRJ_PATH=`pwd`"))
	}

	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide file name!\n")
	}

	conf, err := config.NewConfigFromYAMLPath(prjPath + "/" + os.Args[1])
	if err != nil {
		panic(err)
	}
	log.Println(conf)

	service, err := initService(conf)
	if err != nil {
		panic(err)
	}

	err = service.Start()
	if err != nil {
		log.Fatalln("failed -> ", err)
	}
}
