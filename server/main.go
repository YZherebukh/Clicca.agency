package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Clicca.agency/server/handlers"
	blockchain "github.com/Clicca.agency/server/internal/block_chain"
	"github.com/Clicca.agency/server/proof_of_work"
	"github.com/Clicca.agency/server/quote"
	"github.com/monzo/typhon"
)

func main() {
	log.Println("STARTING SERVER ...")

	chain := blockchain.New()

	difficulty, err := strconv.Atoi(os.Getenv("DIFFICULTY"))
	if err != nil {
		log.Fatalf("missing value for env DIFFICULTY ")
	}

	pow := proof_of_work.New(chain, difficulty)

	router := handlers.NewRoutes(new(typhon.Router), pow, chain, quote.New())
	router.WithRoutes()
	svc := router.Serve().
		Filter(typhon.ErrorFilter).
		Filter(typhon.H2cFilter)

	srv, err := typhon.Listen(svc, ":8080", typhon.WithTimeout(typhon.TimeoutOptions{Read: time.Second * 10}))
	if err != nil {
		log.Fatalf("server failed %s,", err)
	}

	log.Printf("👋  Listening on %v", srv.Listener().Addr())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("☠️  Shutting down")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)

}
