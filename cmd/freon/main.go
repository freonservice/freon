package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/freonservice/freon/internal/config"

	_ "github.com/lib/pq"
	"github.com/powerman/structlog"
)

var (
	log = structlog.New(structlog.KeyUnit, "main")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	config.Init()

	ctxShutdown, shutdown := context.WithCancel(context.Background())
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)

	svc := &service{}
	if err := svc.runServe(ctxShutdown, shutdown); err != nil {
		log.Fatal(err)
	}

	<-sigc
	log.Println("Graceful stop server")
	if err := svc.frontendSrv.Shutdown(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
