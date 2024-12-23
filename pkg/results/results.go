package results

import (
	"fmt"
	"sync"
)

var ResultMap map[string][]int
var ResMutex sync.Mutex

func Init() {
	ResultMap = make(map[string][]int)
}
func ShowResults() {
	for key, value := range ResultMap {
		fmt.Printf("Host %s - %d open ports \n", key, len(value))
	}
}
