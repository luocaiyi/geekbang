// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	// signal process
	group.Go(func() error {
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Start Server Err")
				return ctx.Err()
			case <-sig:
				fmt.Println("Stop Server")
				return nil
			}
		}
	})

	// http server 1
	group.Go(func() error {
		return startServer("localhost:8080", ctx)
	})

	// http server 2
	group.Go(func() error {
		return startServer("localhost:8081", ctx)
	})

	group.Go(func() error {
		time.Sleep(time.Second)
		return errors.New("generate error")
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("Err: %v\n", err)
	}
}

// startServer
func startServer(addr string, ctx context.Context) error {
	server := http.Server{Addr: addr}
	go func() {
		<-ctx.Done()
		fmt.Printf("Server %s Shutdown\n", addr)
		_ = server.Shutdown(ctx)
	}()
	return server.ListenAndServe()
}

// Start Server Err
// Server localhost:8080 Shutdown
// Server localhost:8081 Shutdown
// Err: generate error
