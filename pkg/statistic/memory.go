package statistic

import (
	"context"
	"runtime"
	"sync"
	"time"
)

func MemoryThread(ctx context.Context, mutex *sync.Mutex) {
	var m runtime.MemStats
	for {
		select {
		case <-ctx.Done():
			return
		default:
			runtime.ReadMemStats(&m)
			mutex.Lock()
			AllocatedMemory = m.Alloc
			mutex.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}
}
