package statistic

import (
	"context"
	"time"
)

func BufferThread(ctx context.Context, ScanIPChannel chan string, PingIPChannel chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			StatisticMutex.Lock()

			// Получаем параметры буффера для сканирования
			CheckedLenBuffer = uint64(len(ScanIPChannel))
			CheckedCapBuffer = uint64(cap(ScanIPChannel))

			// Получаем параметры буффера для пингования
			NotCheckedLenBuffer = uint64(len(PingIPChannel))
			NotCheckedCapBuffer = uint64(cap(PingIPChannel))

			StatisticMutex.Unlock()

			time.Sleep(time.Millisecond * 100)
		}
	}
}
