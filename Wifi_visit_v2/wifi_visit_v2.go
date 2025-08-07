package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type TransitionCount map[string]int

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

	transitionCount := make(TransitionCount)
	transitionCount["P10"] = 0
	transitionCount["P01"] = 0
	transitionCount["P1T"] = 0
	transitionCount["P0T"] = 0

	counter := 10000

	fmt.Printf("Parameters:\n")
	fmt.Printf("ET0: %.2f, ET1: %.2f, ETmu: %.2f\n", ET0, ET1, ETmu)
	fmt.Printf("P10 (1→0): %.4f\n", P10)
	fmt.Printf("P01 (0→1): %.4f\n", P01)
	fmt.Printf("P1T (1→T): %.4f\n", P1T)
	fmt.Printf("P0T (0→T): %.4f\n", P0T)
	fmt.Println()

	n1 := runSimulations(counter, ET0, ET1, P10, P01, P1T, P0T, transitionCount)

	displayProbabilities(n1, counter, lambda0, lambda1, mu)

	totalTransitions := transitionCount["P10"] + transitionCount["P01"] + transitionCount["P1T"] + transitionCount["P0T"]
	fmt.Printf("\nTransition Counts:\n")
	fmt.Printf("P10 (1→0): %d\n", transitionCount["P10"])
	fmt.Printf("P01 (0→1): %d\n", transitionCount["P01"])
	fmt.Printf("P1T (1→T): %d\n", transitionCount["P1T"])
	fmt.Printf("P0T (0→T): %d\n", transitionCount["P0T"])
	fmt.Printf("Total transitions: %d\n", totalTransitions)

	fmt.Printf("\nEmpirical Transition Probabilities:\n")
	from1 := transitionCount["P10"] + transitionCount["P1T"]
	if from1 > 0 {
		p10_emp := float64(transitionCount["P10"]) / float64(from1)
		p1T_emp := float64(transitionCount["P1T"]) / float64(from1)
		fmt.Printf("P(1→0): %.6f (expected: %.4f)\n", p10_emp, P10)
		fmt.Printf("P(1→T): %.6f (expected: %.4f)\n", p1T_emp, P1T)
	}

	from0 := transitionCount["P01"] + transitionCount["P0T"]
	if from0 > 0 {
		p01_emp := float64(transitionCount["P01"]) / float64(from0)
		p0T_emp := float64(transitionCount["P0T"]) / float64(from0)
		fmt.Printf("P(0→1): %.6f (expected: %.4f)\n", p01_emp, P01)
		fmt.Printf("P(0→T): %.6f (expected: %.4f)\n", p0T_emp, P0T)
	}
}

func runSimulations(
	counter int,
	ET0, ET1, P10, P01, P1T, P0T float64,
	transitionCount TransitionCount,
) map[int]int {
	freq := make(map[int]int)

	for i := 0; i < counter; i++ {
		initial := initialstate(ET0, ET1)
		countOfOnes, path := wifi_visit(initial, P10, P01, P1T, P0T, transitionCount)
		freq[countOfOnes]++
		// If you want to see the Wifi visit, comment the next line
		_ = path
		// If you want to see the Wifi visit, uncomment the next line
		// fmt.Printf("Iteration %d: Path = %v, n1 = %d\n", i+1, path, countOfOnes)
	}

	return freq
}

func displayProbabilities(n1 map[int]int, total int, lambda0, lambda1, mu float64) {
	SumPn1 := 0.0

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

	fmt.Printf("\n")

	for _, n := range keys {
		fmt.Printf("P(n = %d) = %.6f (case 1)\n", n, case1(lambda0, lambda1, mu, n))
		fmt.Printf("P(n = %d) = %.6f (case 2)\n", n, case2(lambda0, lambda1, mu, n))
		fmt.Printf("P(n = %d) = %.6f (case 3)\n", n, case3(lambda0, lambda1, mu, n))
		fmt.Printf("P(n = %d) = %.6f (case 4)\n", n, case4(lambda0, lambda1, mu, n))
		SumPn1 += case1(lambda0, lambda1, mu, n)
		SumPn1 += case2(lambda0, lambda1, mu, n)
		SumPn1 += case3(lambda0, lambda1, mu, n)
		SumPn1 += case4(lambda0, lambda1, mu, n)

	}

	fmt.Printf("Sum P(n) = %.6f\n", SumPn1)
}

func initialstate(ET0, ET1 float64) int {
	U := rand.Float64()
	P0 := ET0 / (ET0 + ET1)
	if U <= P0 {
		return 0
	}
	return 1
}

func wifi_visit(
	initialstate int,
	P10, P01, P1T, P0T float64,
	transitionCount TransitionCount,
) (int, []int) {
	state := initialstate
	countOfOnes := 0
	_ = P1T
	_ = P0T
	var wifiVisit []int

	wifiVisit = append(wifiVisit, state)
	if state == 1 {
		countOfOnes++
	}

	for {
		if state == 0 {
			U := rand.Float64()
			if U <= P01 {
				state = 1
				wifiVisit = append(wifiVisit, state)
				transitionCount["P01"]++
				countOfOnes++
			} else {
				transitionCount["P0T"]++
				break
			}
		} else if state == 1 {
			U := rand.Float64()
			if U <= P10 {
				state = 0
				wifiVisit = append(wifiVisit, state)
				transitionCount["P10"]++
			} else {
				transitionCount["P1T"]++
				break
			}
		}
	}

	return countOfOnes, wifiVisit
}

func case1(lambda0, lambda1, mu float64, n int) float64 {
	if n == 0 {
		return 0.0
	}
	return (lambda0 / (lambda1 + lambda0)) * math.Pow((math.Pow(mu, 2)/((lambda0+mu)*(lambda1+mu))), float64(n-1)) * (lambda0 / (lambda0 + mu))
}

func case2(lambda0, lambda1, mu float64, n int) float64 {
	if n == 0 {
		return 0.0
	}
	return (lambda0 / (lambda1 + lambda0)) * math.Pow((mu/(lambda0+mu)), float64(n)) * math.Pow((mu/(lambda1+mu)), float64(n-1)) * (lambda1 / (lambda1 + mu))
}

func case3(lambda0, lambda1, mu float64, n int) float64 {
	if n == 0 {
		return 0.0
	}
	return (lambda1 / (lambda1 + lambda0)) * math.Pow((mu/(lambda1+mu)), float64(n)) * math.Pow((mu/(lambda0+mu)), float64(n-1)) * (lambda0 / (lambda0 + mu))
}

func case4(lambda0, lambda1, mu float64, n int) float64 {
	return (lambda1 / (lambda1 + lambda0)) * math.Pow((math.Pow(mu, 2)/((lambda1+mu)*(lambda0+mu))), float64(n)) * (lambda1 / (lambda1 + mu))
}
