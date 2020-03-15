package main

import (
	"github.com/aibotsoft/service-kit/internal/config"
	"github.com/aibotsoft/service-kit/internal/logging"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log := logging.New(cfg)
	log.Println("Beginning...")
	for {
		log.Println(time.Now().UTC())
		time.Sleep(2 * time.Second)
	}
}
