package statistic

import (
	"context"
	"time"
)

func BufferThread(ctx context.Context, IPChannel chan string, RawIPChannel chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			StatisticMutex.Lock()
			NotCheckedLenBuffer = uint64(len(RawIPChannel))
			NotCheckedCapBuffer = uint64(cap(RawIPChannel))

			CheckedLenBuffer = uint64(len(IPChannel))
			CheckedCapBuffer = uint64(cap(IPChannel))
			StatisticMutex.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}
}
