package main

import (
	"github.com/nnqq/scr-classifier/classifier"
	"github.com/nnqq/scr-classifier/classifierimpl"
	"github.com/nnqq/scr-classifier/config"
	"github.com/nnqq/scr-classifier/logger"
	graceful "github.com/nnqq/scr-lib-graceful"
	pbClassifier "github.com/nnqq/scr-proto/codegen/go/classifier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"strings"
	"sync"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	cls, err := classifier.NewClassifier()
	logg.Must(err)

	srv := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
	pbClassifier.RegisterClassifierServer(srv, classifierimpl.NewServer(logg.ZL, cls))

	lis, err := net.Listen("tcp", strings.Join([]string{
		"0.0.0.0",
		cfg.Grpc.Port,
	}, ":"))
	logg.Must(err)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		graceful.HandleSignals(srv.GracefulStop)
	}()
	go func() {
		defer wg.Done()
		logg.Must(srv.Serve(lis))
	}()
	wg.Wait()
}
