package SecondWork

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

//启动http server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/errgroup", HandleRequst)
	fmt.Printf("start httpserver!\n")
	err := srv.ListenAndServe()
	return err
}

//http handler
func HandleRequst(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "errgroup exercise！\n")
}

func Init() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	g, errctx := errgroup.WithContext(ctx)

	srv := &http.Server{Addr: ":8088"}

	g.Go(func() error {
		//路由实现
		return StartHttpServer(srv)
	})

	g.Go(func() error {
		<-errctx.Done()
		fmt.Printf("httpserver stopped!\n")
		return srv.Shutdown(errctx)
	})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)

	g.Go(func() error {
		for {
			select {
			case <-errctx.Done():
				return errctx.Err()
			case <-signalChan:
				cancel()
			}
		}
	})
	err := g.Wait()

	fmt.Println(err)

	fmt.Printf("all group down completed!\n")

}
