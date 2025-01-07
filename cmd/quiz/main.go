package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/praveenmahasena/goquiz/internal"
)

func main() {
	ctx,cancel:=context.WithCancel(context.Background())
	go func() {
		if err := internal.Run(ctx); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1) // pretty sure we don't need os.Exit here
		}
	}()


	sig:=make(chan os.Signal,1)


	signal.Notify(sig,os.Interrupt)
	signal.Notify(sig,syscall.SIGTERM)

	select{
	case <- sig:
		cancel()
	case <-ctx.Done():
		log.Println("shutting down")
	}
}
