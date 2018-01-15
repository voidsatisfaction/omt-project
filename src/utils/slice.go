package utils

import (
	"math/rand"
	"time"
)

type shuffleSlice interface {
	Swap(i, j int)
	Len() int
}

func Shuffle(sfs shuffleSlice) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 1; i < sfs.Len(); i++ {
		randIndex := r.Intn(i)

		sfs.Swap(i, randIndex)
	}
}
