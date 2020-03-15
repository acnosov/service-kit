package logging

import (
	"github.com/aibotsoft/service-kit/internal/config"
	"log"
	"os"
)

func New(cfg *config.Config) *log.Logger {
	prefix := cfg.Service.Name + ": "
	return log.New(os.Stdout, prefix, log.Ldate|log.Lmsgprefix|log.Lmicroseconds|log.Lshortfile)
}
