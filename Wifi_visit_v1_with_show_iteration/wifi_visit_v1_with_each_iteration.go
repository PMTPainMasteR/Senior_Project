package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
	lambda0 := math.Pow(400, -1)
	lambda1 := math.Pow(600, -1)
	mu := math.Pow(400, -1)
	ET0 := 1 / lambda0
	ET1 := 1 / lambda1
	ETmu := 1 / mu
	P10 := mu / (lambda0 + mu)
	P01 := mu / (lambda1 + mu)
	P1T := lambda0 / (lambda0 + mu)
	P0T := lambda1 / (lambda1 + mu)
	counter := 10000

	fmt.Printf("Parameters:\n")
	fmt.Printf("ET0: %.2f, ET1: %.2f, ETmu: %.2f\n", ET0, ET1, ETmu)
	fmt.Printf("P10: %.4f, P01: %.4f, P1T: %.4f, P0T: %.4f\n", P10, P01, P1T, P0T)

	n1 := runSimulations(counter, ET0, ET1, P10, P01, P1T, P0T)

	displayProbabilities(n1, counter)
}

func runSimulations(counter int, ET0, ET1, P10, P01, P1T, P0T float64) map[int]int {
	visit := make(map[int]int)

	for i := 0; i < counter; i++ {
		initial := initialstate(ET0, ET1)
		countOfOnes, wifiVisit := wifi_visit(initial, P10, P01, P1T, P0T)
		visit[countOfOnes]++
		fmt.Printf("Iteration %d: WiFi visit = %v, count of ones = %d\n", i+1, wifiVisit, countOfOnes)
	}

	return visit
}

func displayProbabilities(n1 map[int]int, total int) {
	fmt.Printf("\nResults from %d simulations:\n", total)
	fmt.Printf("Frequency count: %v\n", n1)
	fmt.Printf("\nProbabilities:\n")

	keys := make([]int, 0, len(n1))
	for k := range n1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	totalProbability := 0.0
	for _, k := range keys {
		freq := n1[k]
		probability := float64(freq) / float64(total)
		totalProbability += probability
		fmt.Printf("P(n1 = %d) = %d/%d = %.6f\n", k, freq, total, probability)
	}

	fmt.Printf("\nTotal probability: %.6f\n", totalProbability)
}

func initialstate(ET0, ET1 float64) int {
	U := rand.Float64()
	P0 := ET0 / (ET0 + ET1)
	if U <= P0 {
		return 0
	} else {
		return 1
	}
}

func wifi_visit(initialstate int, P10 float64, P01 float64, P1T float64, P0T float64) (int, []int) {
	state := initialstate
	countOfOnes := 0
	var Wifi_Visit []int

	Wifi_Visit = append(Wifi_Visit, state)
	if state == 1 {
		countOfOnes++
	}

	for {
		if state == 0 {
			U := rand.Float64()
			if U <= P01 {
				state = 1
				Wifi_Visit = append(Wifi_Visit, state)
				countOfOnes++
			} else {
				return countOfOnes, Wifi_Visit
			}
		} else if state == 1 {
			U := rand.Float64()
			if U <= P10 {
				state = 0
				Wifi_Visit = append(Wifi_Visit, state)
			} else {
				return countOfOnes, Wifi_Visit
			}
		}
	}
}
