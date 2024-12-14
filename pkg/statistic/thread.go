package statistic

import (
	"context"
	"fmt"
	"github.com/gosuri/uilive"
	"log"
	"sync"
	"time"
)

var (
	Speed uint64 // Скорость в портах в секунду

	AllocatedMemory uint64 // Сколько выделено памяти

	NotCheckedLenBuffer uint64 // Непроверенные адреса заполненность
	NotCheckedCapBuffer uint64 // Непроверенные адреса вместимость

	CheckedLenBuffer uint64 // Проверенные адреса заполненность
	CheckedCapBuffer uint64 // Проверенные адреса вместимость

	PortsCounter int // Счетчик отсканированных портов
	PortsSpeed   int

	PingStatus uint64
	Mutex      sync.Mutex
)

func StatisticStart(ctx context.Context, IPChannel chan string, RawIPChannel chan string) {
	log.Println("Подсистема статистики запускается")
	go statisticThread(ctx, &Mutex)
	go speedThread(ctx, &Mutex)
	go MemoryThread(ctx, &Mutex)
	go BufferThread(ctx, IPChannel, RawIPChannel, &Mutex)
	go pingThread(ctx, &Mutex)
	log.Println("Подсистема статистики запущена")

	time.Sleep(time.Second)
}

func statisticThread(ctx context.Context, mutex *sync.Mutex) {
	writer := uilive.New()
	writer.Start()
	for {
		select {
		case <-ctx.Done():
			writer.Stop()
			return
		default:
			mutex.Lock()
			fmt.Fprintf(writer, "Scanned ports: %d \n", PortsCounter)
			fmt.Fprintf(writer.Newline(), "Speed: %d ports/sec\n", PortsSpeed)
			fmt.Fprintf(writer.Newline(), "Allocated Memory: %d kB \n", AllocatedMemory/1024)
			fmt.Fprintf(writer.Newline(), "Not Checked Buffer: %d %d \n", NotCheckedLenBuffer, NotCheckedCapBuffer)
			fmt.Fprintf(writer.Newline(), "Checked Buffer: %d %d \n", CheckedLenBuffer, CheckedCapBuffer)
			fmt.Fprintf(writer.Newline(), "Ping status: %d ms \n", PingStatus)
			mutex.Unlock()
			time.Sleep(time.Second / 2)
		}
	}
}
