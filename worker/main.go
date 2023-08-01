package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Clicca.agency/worker/client"
	"github.com/Clicca.agency/worker/internal/executor"
	"github.com/Clicca.agency/worker/internal/task"
)

func main() {
	ctx := context.Background()

	log.Println("starting client ...")

	serverBaseURL := os.Getenv("SERVER_URL")
	if serverBaseURL == "" {
		log.Fatal("failed to get serverBaseURL")
	}

	difficulty, err := strconv.Atoi(os.Getenv("DIFFICULTY"))
	if err != nil {
		log.Fatal("missing value for env DIFFICULTY ", err)
	}

	interval, err := strconv.Atoi(os.Getenv("TIME_INTERVAL"))
	if err != nil {
		log.Fatal("missing value for env TIME_INTERVAL ", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	executor := executor.New(&client.Client{URL: serverBaseURL}, task.NewBlockChain(difficulty))

	for {
		select {
		case <-time.NewTicker(time.Duration(interval * int(time.Second))).C:
			err = executor.Do(ctx)
			if err != nil {
				log.Println(err)
			}
		case <-done:
			log.Println("stopping server ...")
			return
		}
	}
}
