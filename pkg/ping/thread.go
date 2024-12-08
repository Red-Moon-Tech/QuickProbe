package ping

import "sync"

// PingingThread является пингующий горутиной, сначала сюда попадают адреса для проверки

// Создаем WaitGroups для отслеживания завершения горутин
var (
	// PingWG - Группа горутин отвечающих за запуск/завершение пингующих потоков
	PingWG sync.WaitGroup
)

func PingingThread(RawIPChannel chan string, IPChannel chan string) {
	defer PingWG.Done()

	for {
		ip, ok := <-RawIPChannel
		if ok {
			available := pingTest(ip)
			if available {
				IPChannel <- ip
			}
		} else {
			break
		}
	}
}