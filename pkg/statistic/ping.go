package statistic

import (
	"context"
	probing "github.com/prometheus-community/pro-bing"
	"sync"
	"time"
)

func pingThread(ctx context.Context, mutex *sync.Mutex) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			pinger, err := probing.NewPinger("8.8.8.8")
			if err != nil {
				panic(err)
			}

			pinger.Count = 1

			err = pinger.Run()
			if err != nil {
				panic(err)
			}

			stats := pinger.Statistics()

			mutex.Lock()
			if stats.PacketsRecv != 0 {
				PingStatus = uint64(stats.MaxRtt.Milliseconds())
			} else {
				PingStatus = 0
			}
			mutex.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}
}
