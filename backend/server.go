package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler:      mux,
			ReadTimeout:  20 * time.Second,  // リクエスト全体読み込み最大時間
			WriteTimeout: 60 * time.Second,  // レスポンス書き込みの最大時間
			IdleTimeout:  120 * time.Second, // 次リクエストまでの最大待機時間(キープアライブ)
		},
		l: l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// ErrServerClosedは正常にサーバーがシャットダウンされたことを示すため、エラー判定から除外
		if err := s.srv.Serve(s.l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
		s.srv.Close()
	}

	return eg.Wait()
}
