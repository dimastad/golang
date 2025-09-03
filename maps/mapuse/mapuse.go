package mapuse

import (
	"fmt"
	"maps"
)

func CountLanguages(users map[string]string) map[string]int {
	result := map[string]int{}
	values := maps.Values(users)
	fmt.Println("values ", values)
	for val := range values {
		fmt.Println("key", val)
		fmt.Println("res", result)
		result[val]++
	}
	fmt.Println(result)
	return result
}
