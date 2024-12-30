package statistic

import (
	"QuickProbe/pkg/argflags"
	"context"
	"github.com/shirou/gopsutil/net"
	"time"
)

func InterfacesThread(ctx context.Context) {
	stats, _ := net.IOCounters(true)

	for _, stat := range stats {
		select {
		case <-ctx.Done():
			return
		default:
			if *argflags.ShowInterfaceInfo == stat.Name {
				StatisticMutex.Lock()
				IfaceWorkload = stat.BytesSent + stat.BytesRecv
				StatisticMutex.Unlock()
			}
			time.Sleep(time.Second)
		}
	}
}
