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

func StatisticStart(ctx context.Context, ScanIPChannel chan string, PingIPChannel chan string) {
	go statisticThread(ctx)
	go speedThread(ctx)
	go MemoryThread(ctx)
	go BufferThread(ctx, ScanIPChannel, PingIPChannel)
	go pingThread(ctx)
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
			// Форматируем информацию перед выводом
			StatisticMutex.Lock()

			output := ""
			output += fmt.Sprintf("Scanned ports: %d \n", PortsCounter)
			output += fmt.Sprintf("Speed: %d ports/sec\n", PortsSpeed)
			output += fmt.Sprintf("Allocated Memory: %d kB \n", AllocatedMemory/1024)
			output += fmt.Sprintf("Scan Buffer: %d/%d \n", NotCheckedLenBuffer, NotCheckedCapBuffer)
			output += fmt.Sprintf("Ping Buffer: %d/%d \n", CheckedLenBuffer, CheckedCapBuffer)
			output += fmt.Sprintf("Ping status: %d ms \n", PingStatus)

			StatisticMutex.Unlock()

			// Выводим информацию
			fmt.Fprintf(writer, output)

			time.Sleep(time.Second / 2)
		}
	}
}
