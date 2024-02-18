package utils

import "fmt"

func ShowArray(array []int) {
	for i, e := range array {
		fmt.Printf("index %d has the value %d\n", i, e)
	}
}
