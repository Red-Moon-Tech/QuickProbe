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
		fmt.Println("\nHost", key)
		for _, i := range value {
			fmt.Printf("Discovered open port %d on %s\n", i, key)
		}
	}
}
