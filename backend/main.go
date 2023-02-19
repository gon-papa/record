package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gon-papa/record/config"
	"github.com/gon-papa/record/router"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("server not run: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cnf, err := config.GetConfig()
	os.Setenv("TZ", cnf.TimeZone)

	if err != nil {
		fmt.Printf("server not run: %v\n", err)
		os.Exit(1)
	}

	l, err := net.Listen("tcp", cnf.Port)
	if err != nil {
		log.Fatalf("指定ポートでのサーバーの起動に失敗しました。 port%s->error: %v", cnf.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux := router.NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
