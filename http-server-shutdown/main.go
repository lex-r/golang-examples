package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	httpServer := &http.Server{Addr: ":8080", Handler: nil}
	mux := http.NewServeMux()
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 30)
		fmt.Fprint(w, "Hello after 30 seconds!")
	})
	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 1)
		fmt.Fprint(w, "Hello after 1 second!")
	})
	httpServer.Handler = mux

	go httpServer.ListenAndServe()

	s := <-c
	log.Printf("got signal: %s\n", s)

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	err := httpServer.Shutdown(ctx)
	log.Printf("http server stopped, err: %v\n", err)
}
