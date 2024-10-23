package utils

import (
	"math/rand/v2"
)

func GenerateNumbers(totalNumbers int) ([]int, int) {
	var numbers []int
	sum := 0
	
	for sum < 10 {
		//Generates 'n' numbers
		for i:= 0; i < totalNumbers; i++ {
			num := rand.IntN(50) + 1
			numbers = append(numbers, num)
		}
		
		//Generates random sum from numbers
		for i:= 0; i < rand.IntN(totalNumbers - 2) + 2; i++ {
			sum += numbers[i]
		}
		
		if sum < 10 {
			sum = 0
			numbers = nil
		}
	}
	
	//Remove one solution variants
	for i:= 0; i < totalNumbers; i++ {
		if sum == numbers[i] {
			numbers[i] = rand.IntN(50) + 1
			i = i - 1
		}
	}
	
	return shuffle(numbers), sum
}

//
func remove(s []int, i int) []int {
	s[i] = s[len(s) - 1]
	return s[:len(s) - 1]
}

//Shuffle slice
func shuffle(s []int) []int {
	for i := range s {
		j := rand.IntN(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func isOneSolution(s []int, a int) bool {
	for i:= 0; i < len(s); i++ {
		if s[i] == a {
			return true
		}
	}
	return false
}
