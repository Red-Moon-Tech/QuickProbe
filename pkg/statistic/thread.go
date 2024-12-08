package statistic

import (
	"context"
	"fmt"
	"github.com/gosuri/uilive"
	probing "github.com/prometheus-community/pro-bing"
	"log"
	"runtime"
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
)

func StatisticStart(ctx context.Context, IPChannel chan string, RawIPChannel chan string) {
	log.Println("Подсистема статистики запускается")
	var mutex sync.Mutex
	go statisticThread(ctx, &mutex)
	go speedThread(ctx, &mutex)
	go MemoryThread(ctx, &mutex)
	go BufferThread(ctx, IPChannel, RawIPChannel, &mutex)
	go pingThread(ctx, &mutex)
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

func speedThread(ctx context.Context, mutex *sync.Mutex) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mutex.Lock()
			p := PortsCounter
			mutex.Unlock()

			time.Sleep(time.Second)

			mutex.Lock()
			PortsSpeed = PortsCounter - p // Находим скорость как разность портов за секунду
			mutex.Unlock()

		}
	}
}
