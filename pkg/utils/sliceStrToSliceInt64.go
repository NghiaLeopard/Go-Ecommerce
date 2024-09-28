package utils

import (
	"fmt"
	"strconv"
	"sync"
)

func ParseSliceInt64(wg sync.WaitGroup, arr []string) ([]int64, error) {
	newArr := make([]int64, len(arr))
	for i, v := range arr {
		result, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return []int64{}, fmt.Errorf("server Error when parse slice")
		}
		newArr[i] = result
	}

	wg.Done()

	return newArr, nil
}
