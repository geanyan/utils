package main

import (
	"fmt"
	"github.com/laohanlinux/memCache"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	s, _ := memCache.NewNonBlockingChan()

	var g sync.WaitGroup
	currentNumbers, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < currentNumbers; i++ {
		go func() {
			g.Add()
			defer g.Done()
			<-s
			time.Sleep(time.Second * 1)
		}()
	}

	g.Wait()
}