package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ringsaturn/go-web-template/pkg/config"
	pkggrpc "github.com/ringsaturn/go-web-template/pkg/server/grpc"
	pkghttp "github.com/ringsaturn/go-web-template/pkg/server/http"
	"golang.org/x/sync/errgroup"
)

var ErrServerNotReady = errors.New("server is nil")

type Service struct {
	conf       *config.Config
	HTTPServer *pkghttp.Server
	GRPCServer *pkggrpc.Server
}

func NewService(conf *config.Config, httpServer *pkghttp.Server, gRPCServer *pkggrpc.Server) (*Service, error) {
	return &Service{
		conf:       conf,
		HTTPServer: httpServer,
		GRPCServer: gRPCServer,
	}, nil
}

func (s *Service) StartHTTP(ctx context.Context) error {
	if s.HTTPServer == nil {
		return ErrServerNotReady
	}
	go func() {
		<-ctx.Done()
		log.Println("http ctx done")
		_ = s.HTTPServer.Shutdown(ctx)
	}()
	return s.HTTPServer.Start(ctx)
}

func (s *Service) StartGRPC(ctx context.Context) error {
	if s.GRPCServer == nil {
		return ErrServerNotReady
	}
	go func() {
		<-ctx.Done()
		log.Println("http ctx done")
		_ = s.GRPCServer.Shutdown(ctx)
	}()
	return s.GRPCServer.Start(ctx)
}

func (s *Service) StartPProfAPI(ctx context.Context) error {
	app := &http.Server{
		Addr:           s.conf.PProfServer.Addr,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		<-ctx.Done()
		_ = app.Shutdown(ctx)
	}()
	return app.ListenAndServe()
}

func (s *Service) Start() error {
	ctx, done := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.StartHTTP(gctx)
	})

	g.Go(func() error {
		return s.StartGRPC(gctx)
	})

	g.Go(func() error {
		return s.StartPProfAPI(gctx)
	})

	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-gctx.Done():
				fmt.Println("signal ctx done")
				return gctx.Err()
			case <-sig:
				log.Println("get sig")
				done()
				return nil
			}
		}
	})

	return g.Wait()

}
