package statistic

import (
	"context"
	"sync"
	"time"
)

func BufferThread(ctx context.Context, IPChannel chan string, RawIPChannel chan string, mutex *sync.Mutex) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mutex.Lock()
			NotCheckedLenBuffer = uint64(len(RawIPChannel))
			NotCheckedCapBuffer = uint64(cap(RawIPChannel))

			CheckedLenBuffer = uint64(len(IPChannel))
			CheckedCapBuffer = uint64(cap(IPChannel))
			mutex.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}
}
