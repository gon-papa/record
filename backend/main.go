package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gon-papa/record/router"
)

func init() {
	os.Setenv("TZ", "Asia/Tokyo")
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("server not run: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("指定ポートでのサーバーの起動に失敗しました。 port:8080->error: %v", err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux := router.NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
