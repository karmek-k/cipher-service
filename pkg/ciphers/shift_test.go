package ciphers_test

import (
	"testing"

	"github.com/karmek-k/cipher-service/pkg/ciphers"
)

func Test_ShiftRune_NoLoop(t *testing.T) {
	makeShiftRuneTest(t, 'a', 2, 'c')
	makeShiftRuneTest(t, 'f', -1, 'e')
	makeShiftRuneTest(t, 'f', 0, 'f')
}

func Test_ShiftRune_Loop(t *testing.T) {
	makeShiftRuneTest(t, 'a', -1, 'z')
	makeShiftRuneTest(t, 'z', 1, 'a')
}

func makeShiftRuneTest(t *testing.T, input rune, shift int, expected rune) {
	result, err := ciphers.ShiftRune(input, shift)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if *result != expected {
		t.Errorf("ShiftRune('%c', %v) == '%c', expected '%c'", input, shift, *result, expected)
	}
}
