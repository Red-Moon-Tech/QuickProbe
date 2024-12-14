package statistic

import (
	"context"
	"fmt"
	"github.com/gosuri/uilive"
	"sync"
	"time"
)

var (
	AllocatedMemory uint64 // Сколько выделено памяти

	NotCheckedLenBuffer uint64 // Непроверенные адреса заполненность
	NotCheckedCapBuffer uint64 // Непроверенные адреса вместимость

	CheckedLenBuffer uint64 // Проверенные адреса заполненность
	CheckedCapBuffer uint64 // Проверенные адреса вместимость

	PortsCounter int // Счетчик отсканированных портов
	PortsSpeed   int // Скорость в портах в секунду

	PingStatus     uint64
	StatisticMutex sync.Mutex
)

func StatisticStart(ctx context.Context, IPChannel chan string, RawIPChannel chan string) {
	go statisticThread(ctx)
	go speedThread(ctx)
	go MemoryThread(ctx)
	go BufferThread(ctx, IPChannel, RawIPChannel)
	go pingThread(ctx)

	time.Sleep(time.Second)
}

func statisticThread(ctx context.Context) {
	writer := uilive.New()
	writer.Start()
	for {
		select {
		case <-ctx.Done():
			writer.Stop()
			return
		default:
			StatisticMutex.Lock()
			fmt.Fprintf(writer, "Scanned ports: %d \n", PortsCounter)
			fmt.Fprintf(writer.Newline(), "Speed: %d ports/sec\n", PortsSpeed)
			fmt.Fprintf(writer.Newline(), "Allocated Memory: %d kB \n", AllocatedMemory/1024)
			fmt.Fprintf(writer.Newline(), "Not Checked Buffer: %d %d \n", NotCheckedLenBuffer, NotCheckedCapBuffer)
			fmt.Fprintf(writer.Newline(), "Checked Buffer: %d %d \n", CheckedLenBuffer, CheckedCapBuffer)
			fmt.Fprintf(writer.Newline(), "Ping status: %d ms \n", PingStatus)
			StatisticMutex.Unlock()
			time.Sleep(time.Second / 2)
		}
	}
}
