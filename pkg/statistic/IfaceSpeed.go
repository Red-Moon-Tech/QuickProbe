package statistic

import (
	"context"
	"github.com/showwin/speedtest-go/speedtest"
	"time"
)

func IfaceSpeed(ctx context.Context) {
	var speedtestClient = speedtest.New()
	serverList, _ := speedtestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})
	s := targets[0]
	for _, server := range targets {
		server.PingTest(nil)
		if server.Latency < s.Latency {
			s = server
		}
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			s.DownloadTest()
			s.UploadTest()
			StatisticMutex.Lock()
			IfaceInSpeed = s.DLSpeed
			IfaceOutSpeed = s.ULSpeed
			StatisticMutex.Unlock()

			time.Sleep(time.Second)
		}
	}
}
