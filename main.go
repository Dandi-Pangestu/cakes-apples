package main

import "fmt"

func main() {
	countBoxes, countEachBox := boxing(20, 25)
	fmt.Printf("Count box: %d box, Count each box: %d cakes and apples", countBoxes, countEachBox)
}

func boxing(cakes int, apples int) (int, int) {
	return countBoxes(cakes, apples), countEachBox(cakes, apples)
}

func countBoxes(cakes int, apples int) int {
	if cakes == apples {
		return cakes / getModulus(cakes)
	}

	if cakes < apples {
		return cakes
	} else {
		return apples
	}
}

func countEachBox(cakes int, apples int) int {
	if cakes == apples {
		return getModulus(cakes)
	}

	return 1
}

func getModulus(count int) int {
	for i := 2; i <= 9; i++ {
		if count%i == 0 {
			return i
		}
	}

	return 1
}
