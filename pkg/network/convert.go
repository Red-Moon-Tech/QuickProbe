package network

import (
	"math"
	"strconv"
	"strings"
)

// Преобразование IP адреса в число
func ipToInt(ip string) uint {
	splitIP := strings.Split(ip, ".")
	var intIP uint
	for i := 0; i < 4; i++ {
		cnt, err := strconv.Atoi(splitIP[i])
		if err != nil {
			panic(err)
		}
		intIP += uint(cnt) * uint(math.Pow(2, float64(24-i*8)))
	}

	return intIP
}
