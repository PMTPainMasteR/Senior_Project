package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	counter := 100
	X_m := 196.0
	alpha := 50.0
	var paretoSlice = make([]float64, counter)
	for i := 0; i < counter; i++ {
		value := generatePareto(alpha, X_m)
		paretoSlice[i] = value
		fmt.Printf("Generated Pareto value %d: %f\n", i+1, value)
	}
	arrayCheck(paretoSlice, X_m)
	fmt.Println("Sum of generated Pareto values:", sumArray(paretoSlice)/float64(counter))
}

func generatePareto(alpha, x_m float64) float64 {
	U := rand.Float64()
	if U == 0 {
		for U == 0 {
			U = rand.Float64()
		}
	}
	return x_m / math.Pow(U, 1/alpha)
}

func sumArray(arr []float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func arrayCheck(arr []float64, X_m float64) {
	for i, value := range arr {
		if value < X_m {
			fmt.Printf("Value found at index %d: %f\n", i, value)
		}
	}
}
