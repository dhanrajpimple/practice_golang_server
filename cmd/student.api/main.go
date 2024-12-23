package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dhanrajpimple/practice_golang_server/internal/config"
)

func main(){
    cfg := config.MustLoad()


   router := http.NewServeMux()

   router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, world!"))
   })

   server := http.Server{
	Addr : cfg.Addr,
	Handler : router,

   }
    slog.Info("server is running")

   fmt.Printf("server is running %s", cfg.HTTPServer.Addr)
  
    done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
  
	go func(){
		err := server.ListenAndServe()

		if err != nil {
		 log.Fatal("Fail to start server")
		}

	}()
   
  <-done

 slog.Info("shutting done the server")
  
 ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
 defer cancel()

err :=server.Shutdown(ctx)

 if err != nil {
	slog.Error("failed to shutdown server", slog.String("error", err.Error()))
 }


 slog.Info("server shutdown successfully")

}