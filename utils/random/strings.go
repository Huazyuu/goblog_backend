package random

import "math/rand"

var letters = []rune("abcdefghijklmnopqresuvwxyzABCDEFGHIJKLMNOPQRESUVWXYZ")

func Strings(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
