package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	counter := 10000
	mean := 200
	x := 0.0
	y := 0.0
	var exponentialArray [10000]float64
	for i := 0; i < counter; i++ {
		value := generateExponential(mean)
		x += value
		exponentialArray[i] = value
		y += rand.ExpFloat64() * float64(mean)
	}
	fmt.Println(x / float64(counter))
	fmt.Println("Mean of generated exponential values:", sumArray(exponentialArray)/float64(counter))
	fmt.Println("Mean of random exponential values:", y/float64(counter))
}

func generateExponential(mean int) float64 {
	U := rand.Float64()
	if U == 0 {
		for U == 0 {
			U = rand.Float64()
		}
	}
	return -1 * float64(mean) * math.Log(1-U)
}

func sumArray(arr [10000]float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum
}
