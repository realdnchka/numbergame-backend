package utils

import (
	"strconv"
	"testing"
)

var arr = []int{3, 4, 5, 6, 7, 8, 9, 10}

func TestGenerateNumberCorrectTotalOfNumbers(t *testing.T) {
	for _, count := range arr {
		t.Run("Generates correct total of numbers", func(t *testing.T) {
			numbers, _ := GenerateNumbers(count)

			if len(numbers) != count {
				t.Errorf("Total of numbers is incorrect. Got: %d, expected: %d", len(numbers), count)
			}
		})
	}
}

func TestGenerateNumberMoreThanOneSolution(t *testing.T) {
	for _, count := range arr {
		t.Run("Not generates one-num solution", func(t *testing.T) {
			numbers, sum := GenerateNumbers(count)

			for i := 0; i < len(numbers); i++ {
				if numbers[i] == sum {
					t.Errorf("\nOne-num solution discover! Number is: %d, index is: %d, slice is: %ssummary is: %d", numbers[i], i, convertArrayIntToString(numbers), sum)
				}
			}
		})
	}
}

func TestGenerateNumbersSumValidation(t *testing.T) {
	for _, count := range arr {
		t.Run("Sum validation", func(t *testing.T) {
			numbers, sum := GenerateNumbers(count)

			if !canSumSubset(GenerateNumbers(count)) {
				t.Errorf("\nWanted %d, numbers are: %s", sum, convertArrayIntToString(numbers))
			}
		})
	}
}

func TestGenerateNumbersAnyNumberLessThanSum(t *testing.T) {
	for _, count := range arr {
		t.Run("Any number less than sum validation", func(t *testing.T) {
			numbers, sum := GenerateNumbers(count)

			for n := range numbers {
				if n > sum {
					t.Errorf("\nNumber: %d is greater than sum: %d", n, sum)
				}
			}
		})
	}
}

func convertArrayIntToString(arr []int) string {
	output := ""
	for _, v := range arr {
		temp := strconv.Itoa(v)
		output = output + temp + ", "
	}

	return output
}

// Helper function: Subset sum check
func canSumSubset(numbers []int, sum int) bool {
	s := map[int]struct{}{0: {}} // Создаем множество с начальным значением 0

	for _, n := range numbers {
		if n <= sum {
			b := make([]int, 0) // Создаем временный срез для новых сумм
			for x := range s {
				if x+n <= sum {
					b = append(b, x+n) // Добавляем новую сумму, если она <= m
				}
			}
			for _, sum := range b {
				s[sum] = struct{}{} // Обновляем множество
			}
		}
	}

	_, exists := s[sum] // Проверяем, существует ли сумма m в множестве
	return exists     // Return true if we can form the target sum
}
