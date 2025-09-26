package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temperature := range temperatures {
		var key int
		if temperature >= 0 {
			// для положительных просто floor
			key = int(math.Floor(temperature/10.0)) * 10

		} else {
			// для отрицательных округляем "к нулю"
			key = int(math.Ceil(temperature/10.0)) * 10
		}
		groups[key] = append(groups[key], temperature)
	}

	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%3d: %v\n", k, groups[k])
	}
}
