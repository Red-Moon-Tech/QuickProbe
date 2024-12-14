package statistic

import (
	"context"
	"sync"
	"time"
)

func speedThread(ctx context.Context, mutex *sync.Mutex) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mutex.Lock()
			p := PortsCounter
			mutex.Unlock()

			time.Sleep(time.Second)

			mutex.Lock()
			PortsSpeed = PortsCounter - p // Находим скорость как разность портов за секунду
			mutex.Unlock()

		}
	}
}
