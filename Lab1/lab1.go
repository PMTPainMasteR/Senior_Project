package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var randArray [5]int
	counter := 15000
	for i := 0; i < counter; i++ {
		U := rand.Float32()

		if U < 0.4 {
			randArray[0] += 1
		} else if 0.4 < U && U < 0.4+0.25 {
			randArray[1] += 1
		} else if 0.4+0.25 < U && U < 0.4+0.25+0.2 {
			randArray[2] += 1
		} else if 0.4+0.25+0.2 < U && U < 0.4+0.25+0.2+0.05 {
			randArray[3] += 1
		} else {
			randArray[4] += 1
		}
	}
	fmt.Println("Random Array: ", randArray)
	fmt.Printf("P[x=1] : %.2f\n", float32(randArray[0])/float32(counter))
	fmt.Printf("P[x=2] : %.2f\n", float32(randArray[1])/float32(counter))
	fmt.Printf("P[x=3] : %.2f\n", float32(randArray[2])/float32(counter))
	fmt.Printf("P[x=4] : %.2f\n", float32(randArray[3])/float32(counter))
	fmt.Printf("P[x=5] : %.2f\n", float32(randArray[4])/float32(counter))
}
