package ports

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// StringToPortsList - Преобразует необработанную строку с портами в массив
func StringToPortsList(rawPortsList string) []int {
	// Обьявляем массив портов для возвращения
	var PortsArray = make([]int, 0)

	// Получаем список портов путём разделения строки через запятую
	pa := strings.Split(rawPortsList, ",")

	// Преобразовываем строку в порт или их список
	for _, value := range pa {
		switch {
		case strings.Contains(value, "-"):
			// Получаем диапазон портов
			portRange := portRangeToArray(value)

			// Добавляем полученный диапазон портов в основной массив
			PortsArray = append(PortsArray, portRange...)
		case strings.Contains(value, ".json"):
			// Получаем список портов из файла
			portsList := getPortsFromFile(value)

			// Добавляем полученный список портов в основной массив
			PortsArray = append(PortsArray, portsList...)
		default:
			// Преобразуем из строки в число
			port, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			// Добавляем порт в массив
			PortsArray = append(PortsArray, port)
		}
	}

	return PortsArray
}

// Преобразует диапазон портов в массив
func portRangeToArray(portRange string) []int {
	// Разбиваем диапазон на начальный и конечный порт
	ports := strings.Split(portRange, "-")

	// Преобразуем начальный порт в число
	initial, err := strconv.Atoi(ports[0])
	if err != nil {
		panic(err)
	}

	// Преобразуем конечный порт в число
	final, err := strconv.Atoi(ports[1])
	if err != nil {
		panic(err)
	}

	// Создаём массив содержащий диапазон портов
	sp := make([]int, (final-initial)+1)
	for i := 0; i <= final-initial; i++ {
		sp[i] = initial + i
	}

	return sp
}

// Получает список портов из файла
func getPortsFromFile(fileName string) []int {
	// Читаем данные из файла
	content, err := os.ReadFile("/usr/share/QuickProbe/ports_list/" + fileName)
	if err != nil {
		panic(err)
	}

	// Обьявляем структуру JSON файла
	type portListFile struct {
		Ports []int `json:"ports"`
	}

	// Парсим JSON файл
	data := portListFile{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}

	return data.Ports
}
