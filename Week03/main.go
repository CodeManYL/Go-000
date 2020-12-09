package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Server (ctx context.Context,g *errgroup.Group,add string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("服务%s 开始处理请求 \n",add)
		time.Sleep(time.Second*5)
		fmt.Printf("服务%s 结束处理请求 \n",add)

	})
	s := &http.Server{
		Addr:         add,
		Handler:      mux,
	}

	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case <-sigs:
			fmt.Printf("服务%s 收到指令 \n",add)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			s.Shutdown(ctx)
			return nil
		case <- ctx.Done():
			//等5秒处理请求的时间
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			s.Shutdown(ctx)
			//server监听的err在shutdonw之前返回 errgroup只会接受第一次的错误 所以这里返回nil都一样
			return nil
		}
	})

	return s.ListenAndServe()
}



func main (){
	g,ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		//这里传入g是因为处理Shutdown的goroutine不等待，处理未请求完也会立刻关闭。
		return Server(ctx,g,":8081")
	})

	g.Go(func() error {
		return Server(ctx,g,":8082")
	})

	if err := g.Wait(); err != nil {
		//报警等错误处理
		fmt.Println(err)
	}

}
