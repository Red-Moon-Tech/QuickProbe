package argflags

import (
	"QuickProbe/pkg/ports"
	"flag"
	"log"
)

// Инициализируем переменные под хранение флагов
var (
	InputNet          *string // Сеть для сканирования
	SkipAddressRange  *string // Флаг пропуска диапазона адрессов
	SkipPingThreads   *bool   // Флаг пропуска пингующих поток для принудительного сканирования
	NumberScanThreads *uint64 // Количество потокв для сканирования
	NumberPingThreads *uint64 // Количество потокв для пингования
	AddressBufferSize *uint64 // Размер буфера адресов
	Timeout           *uint64 // Таймаут при подключении к хосту (мс)
	SkipPrivateRange  *bool   // Флаг пропуска приватных диапазонов
	rawPortsList      *string // Строка с необработанными портам
	PortsList         []int   // Готовый к использованию массив портов
)

func InitFlags() {
	// Определяем флаги
	InputNet = flag.String("Network", "None", "Сеть для сканирования")
	NumberScanThreads = flag.Uint64("NumberScanThreads", 5, "Количество сканирующих потоков")
	NumberPingThreads = flag.Uint64("NumberPingThreads", 5, "Количество пингующих потоков")
	AddressBufferSize = flag.Uint64("AddressBufferSize", 0, "Размеров буфера адресов")
	Timeout = flag.Uint64("Timeout", 100, "Таймаут при подключении к хосту (мс)")
	SkipPrivateRange = flag.Bool("SkipPrivateRange", true, "Пропуск приватных диапазонов при сканировании")
	rawPortsList = flag.String("Ports", "None", "Список портов для сканирования")
	SkipAddressRange = flag.String("SkipAddressRange", "None", "Пропуск диапозона адрессов")
	SkipPingThreads = flag.Bool("SkipPingThreads", false, "Пропуск пингующих потоков для принудительного сканирования")

	// Определяем алиасы для флагов
	flag.StringVar(InputNet, "n", "None", "Сеть для сканирования (алиас InputNet)")
	flag.Uint64Var(NumberScanThreads, "sT", 5, "Количество сканирующих потоков (алиас NumberScanThreads)")
	flag.Uint64Var(NumberPingThreads, "pT", 5, "Количество пингующих потоков (алиас NumberPingThreads)")
	flag.Uint64Var(AddressBufferSize, "bS", 0, "Размеров буфера адресов (алиас AddressBufferSize)")
	flag.Uint64Var(Timeout, "t", 100, "Таймаут при подключении к хосту (мс) (алиас Timeout)")
	flag.BoolVar(SkipPrivateRange, "sP", true, "Пропуск приватных диапазонов при сканировании (алиас SkipPrivateRange)")
	flag.StringVar(rawPortsList, "p", "None", "Список портов для сканирования (алиас Ports)")
	flag.StringVar(SkipAddressRange, "sA", "None", "Пропуск диапозона адрессов (алиас SkipAddressRange)")
	flag.BoolVar(SkipPingThreads, "sPi", false, "Пропуск пингующих потоков для принудительного сканирования (алиас SkipPingThreads)")
}

func ParseFlags() {
	flag.Parse()

	// Если не задан размер буферов, то устанавливаем по количеству сканирующих потоков
	if *AddressBufferSize == 0 {
		*AddressBufferSize = *NumberScanThreads
	}

	// Парсим список портов
	if *rawPortsList == "None" {
		PortsList = ports.StringToPortsList("1-1024")
	} else {
		PortsList = ports.StringToPortsList(*rawPortsList)
	}
}

func CheckFlags() {
	if *InputNet == "None" {
		log.Fatal("Не указана сеть для сканирования, используйте флаг: --Network")
	}
	if *AddressBufferSize < *NumberScanThreads {
		log.Println("WARNING: Рекомендуется устанавливать размер буфера адресов неменее количества сканирующих потоков")
	}
}
