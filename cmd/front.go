package main

import (
	"github.com/auctionee/front/pkg/logger"
	"github.com/auctionee/front/pkg/server"
	"sync"
)

const DEFAULT_PORT = 8081

func main() {
	wg := sync.WaitGroup{}
	s := server.NewAuthServer(DEFAULT_PORT)
	wg.Add(1)
	go func() {
		logger.GetLogger(s.Ctx).Println("Listening on:", DEFAULT_PORT)
		s.Start()
	}()
	wg.Wait()
}
