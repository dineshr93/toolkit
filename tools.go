package toolkit

import (
	"crypto/rand"
)

type Tools struct {
}

func (t *Tools) Randomstring(sizeofrune int) string {
	randomString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+_"
	empty_rune := make([]rune, sizeofrune)

	rune_slice_of_random_string := []rune(randomString)

	for i := range empty_rune {
		p, _ := rand.Prime(rand.Reader, len(rune_slice_of_random_string))
		x := p.Uint64()
		y := uint64(len(rune_slice_of_random_string))
		empty_rune[i] = rune_slice_of_random_string[x%y]

	}
	return string(empty_rune)

}
