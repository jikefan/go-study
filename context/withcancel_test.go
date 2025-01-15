package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("我闭嘴了")
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
