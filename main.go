package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/benthosdev/benthos/v4/public/service"
	_ "github.com/benthosdev/benthos/v4/public/components/all"
	_ "github.com/sansayin/benthos-kafka/bloblang"
)

func init() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Println("Go Routines: ", runtime.NumGoroutine())
			}
		}
	}()
}

func main() {
	service.RunCLI(context.Background())
}
