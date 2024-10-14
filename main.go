package main

import "fmt"

func main() {

	ints := map[string]int64{
		"first":  32,
		"second": 18,
	}

	floats := map[string]float64{
		"first":  10.5,
		"second": 20.20,
	}
	fmt.Printf("Non-Generic Sums: %v and %v \n",
		SumInts(ints),
		SumFloats(floats))

}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
