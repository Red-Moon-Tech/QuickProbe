package statistic

import (
	"context"
	"runtime"
	"time"
)

func MemoryThread(ctx context.Context) {
	var m runtime.MemStats
	for {
		select {
		case <-ctx.Done():
			return
		default:
			runtime.ReadMemStats(&m)
			StatisticMutex.Lock()
			AllocatedMemory = m.Alloc
			StatisticMutex.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}
}
