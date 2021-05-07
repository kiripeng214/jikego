package xnet

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type KiApp struct {
	opts   options
	ctx    context.Context
	cancel func()
}

//New 初始化一些东西
func New(opts ...Option) *KiApp {
	options := options{
		ctx:     context.Background(),
		signals: []os.Signal{syscall.SIGINT, syscall.SIGTERM},
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &KiApp{
		opts:   options,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (this *KiApp) Run() error {
	g, ctx := errgroup.WithContext(this.ctx)
	g.Go(func() error {
		<-ctx.Done()
		return this.opts.HttpServer.Shutdown(ctx)
	})
	g.Go(func() error {
		err := this.opts.HttpServer.ListenAndServe()
		return err
	})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, this.opts.signals...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-quit:
				this.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (this *KiApp) Stop() error {
	if this.cancel != nil {
		this.cancel()
	}
	return nil
}

//func InitServer() error {
//	srv := initRouter()
//	go func() {
//		srv.ListenAndServe()
//	}()
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	select {
//	case <-quit:
//		fmt.Println("shut down the server")
//	}
//	//关闭其他服务
//	if err := srv.Shutdown(context.TODO()); err != nil {
//		fmt.Println("Server Shutdown: ", err)
//		return err
//	}
//	fmt.Println("Server Shutdown: success")
//	return nil
//}
