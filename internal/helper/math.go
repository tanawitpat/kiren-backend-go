package helper

import "math/rand"

// GenerateRandomInteger generates random integer between 1 to user-defined integer.
func GenerateRandomInteger(length int) []int {
	list := rand.Perm(length)
	for i := range list {
		list[i]++
	}
	return list
}
