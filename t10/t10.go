package main

import "fmt"

func main() {
	degrees := map[int][]float32{}

	ex := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, v := range ex {
		group := (int(v) / 10) * 10
		if degrees[group] == nil {
			degrees[group] = make([]float32, 0, 0)
		}
		degrees[group] = append(degrees[group], v)
	}

	fmt.Println(degrees)
}
