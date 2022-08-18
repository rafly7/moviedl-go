package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"moviedl/configs"
	"moviedl/server"
	"net"
	"net/http"
)

func main() {
	// s := lk21.Download()
	// log.Println(finds)
	host, port, secretContext := configs.ConfigServer()
	address := fmt.Sprintf("%s:%s", host, port)
	mux, handler := server.HandleServer()
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    address,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, secretContext, l.Addr().String())
			return ctx
		},
	}

	go func() {
		serverOne.Handler = handler
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server one closed\n")
		} else if err != nil {
			log.Fatalf("error listening for server : %s\n", err)
		}
		cancelCtx()
	}()
	log.Printf("\n\n\n")
	log.Println(">>>>>>>>> START SERVER >>>>>>>>>")
	log.Printf("Secret context %s\n", secretContext)
	log.Printf("Server listening in %s\n", address)
	<-ctx.Done()
}
