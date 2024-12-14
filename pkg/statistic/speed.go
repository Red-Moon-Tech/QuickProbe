package statistic

import (
	"context"
	"time"
)

func speedThread(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			StatisticMutex.Lock()
			p := PortsCounter
			StatisticMutex.Unlock()

			time.Sleep(time.Second)

			StatisticMutex.Lock()
			PortsSpeed = PortsCounter - p // Находим скорость как разность портов за секунду
			StatisticMutex.Unlock()

		}
	}
}
