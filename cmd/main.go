package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Alitindrawan24/go-books-api/internal/pkg/logger"
	http "github.com/Alitindrawan24/go-books-api/internal/server"
)

func main() {
	ctx := context.Background()
	server := http.NewHTTP(ctx)

	go func() {
		err := server.Run().ListenAndServe()
		if err != nil {
			logger.Error(ctx, nil, err, err.Error())
			log.Fatal("server.Run().ListenAndServe() error - mainHTTP")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Printf("Server stopped\n")

	os.Exit(0)
}
