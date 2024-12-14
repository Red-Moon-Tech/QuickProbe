package scan

import (
	"sync"
)

var (
	// WorkWG - Группа горутин отвечающих за запуск/завершение сканирующих потоков
	WorkWG sync.WaitGroup
)

func ScannerThread(IPChannel chan string) {
	defer WorkWG.Done()
	for {
		ip, ok := <-IPChannel
		if ok {
			ports := scanHost(ip)
			if len(ports) != 0 {
				for _, port := range ports {
					println(ip, port)
				}
			}
		} else {
			break
		}
	}
}
