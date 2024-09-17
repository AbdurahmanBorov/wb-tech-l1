package main

import (
	"context"
	"time"
)

func workerCh(stopChan <-chan struct{}) {
	for {
		select {
		case <-stopChan:
			return
		default:
		}
	}
}

func workerCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	stopChan := make(chan struct{})

	//с помощью канала
	go workerCh(stopChan)
	close(stopChan)

	//с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())
	go workerCtx(ctx)
	cancel()

	//таймер
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go workerCtx(ctx)
	time.Sleep(3 * time.Second)
}
