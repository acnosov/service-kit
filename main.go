package main

import (
	"fmt"
	"github.com/aibotsoft/service-kit/internal/config"
	"github.com/aibotsoft/service-kit/internal/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log := logging.New(cfg)
	log.Println("Beginning...")
	log.Println("Config: ", cfg)

	//for {
	//	log.Println(time.Now().UTC())
	//	time.Sleep(2 * time.Second)
	//}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Fatal(http.ListenAndServe(":80", nil))
	//http.ListenAndServe()
}
