package util

import (
	"math/rand"
)

var taskIDletters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = taskIDletters[rand.Intn(len(taskIDletters))]
	}
	return string(b)
}
