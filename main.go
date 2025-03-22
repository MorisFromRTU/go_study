package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func SumSlice(slice []int) int {
	var sum int
	for _, elem := range slice {
		sum += elem
	}
	return sum
}

func FilterPositive(slice []int) []int {
	resultSlice := []int{}
	for _, num := range slice {
		if num > 0 {
			resultSlice = append(resultSlice, num)
		}
	}
	return resultSlice
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	s = strings.Map(clearRune, s)
	srune := []rune(s)
	sruneLen := len(srune)
	for i := 0; i < sruneLen/2; i++ {
		if srune[i] != srune[sruneLen-i-1] {
			return false
		}
	}
	return true
}

func clearRune(r rune) rune {
	if unicode.IsDigit(r) || unicode.IsLetter(r) {
		return r
	}
	return -1
}

func RotateSlice(slice []int, rotate int) []int {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return slice
	}
	rotate = rotate % len(slice)
	if rotate == 0 {
		return slice
	}

	return append(slice[sliceLen-rotate:sliceLen], slice[0:sliceLen-rotate]...)
}

func GroupByCategory(s string) map[string][]rune {
	// categories := make(map[string][]rune)
	categories := map[string][]rune{
		"letters": {},
		"digits":  {},
		"others":  {},
	}
	for _, char := range s {
		if unicode.IsLetter(char) {
			categories["letters"] = append(categories["letters"], char)
		} else if unicode.IsDigit(char) {
			categories["digits"] = append(categories["digits"], char)
		} else {
			categories["others"] = append(categories["others"], char)
		}
	}
	return categories
}

func CharacterFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char] += 1
	}
	return frequency
}

func CompressString(s string) string {
	var currentChr rune = rune(s[0])
	var currentCount int = 1
	var resultStr strings.Builder
	if len(s) == 0 {
		return s
	}

	for _, char := range s[1:] {
		if char == currentChr {
			currentCount += 1
		} else {
			resultStr.WriteString(string(currentChr))
			resultStr.WriteString(strconv.Itoa(currentCount))
			currentChr = char
			currentCount = 1

		}
	}
	resultStr.WriteString(string(currentChr))
	if currentCount > 1 {
		resultStr.WriteString(strconv.Itoa(currentCount))
	}
	return resultStr.String()
}

func SumDiagonal(matrix [][]int) int {
	n := len(matrix)
	if n == 0 || len(matrix[0]) != n {
		fmt.Println("Матрица не квадратная!")
		return 0
	}

	var sum int
	for i := 0; i < n; i++ {
		sum += matrix[i][i]
	}
	return sum
}

func main() {

}
