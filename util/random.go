package util

import (
	"math/rand"
	"time"
)

var taskIDletters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = taskIDletters[rand.Intn(len(taskIDletters))]
	}
	return string(b)
}
