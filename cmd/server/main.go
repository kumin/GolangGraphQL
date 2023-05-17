package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/kumin/GolangGraphQL/apps/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	server, err := server.BuildServer()
	if err != nil {
		panic(err)
	}
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		errchan := make(chan error)
		go func() {
			if err := server.Start(ctx); err != nil {
				errchan <- err
			} else {
				close(errchan)
			}
		}()
		for {
			select {
			case err := <-errchan:
				return err
			case <-ctx.Done():
				return nil
			}
		}
	})
	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
