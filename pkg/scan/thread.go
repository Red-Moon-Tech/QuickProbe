package scan

import "sync"

var (
	// WorkWG - Группа горутин отвечающих за запуск/завершение сканирующих потоков
	WorkWG sync.WaitGroup
)

func ScannerThread(IPChannel chan string, PortsCount *int) {
	defer WorkWG.Done()
	var mutex sync.Mutex
	for {
		ip, ok := <-IPChannel
		if ok {
			ports := scanHost(ip, PortsCount, &mutex)
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
