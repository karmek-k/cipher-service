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

func Test_DetermineShift_Small(t *testing.T) {
	assertCorrectShift(t, 0, 0)
	assertCorrectShift(t, 1, 1)
	assertCorrectShift(t, 10, 10)
}

func Test_DetermineShift_Large(t *testing.T) {
	assertCorrectShift(t, ciphers.AlphabetLength, 0)
	assertCorrectShift(t, ciphers.AlphabetLength+1, 1)
	assertCorrectShift(t, ciphers.AlphabetLength*10, 0)
	assertCorrectShift(t, 1000, 1000%ciphers.AlphabetLength)
}

func Test_DetermineShift_NegativeSmall(t *testing.T) {
	assertCorrectShift(t, -1, ciphers.AlphabetLength-1)
	assertCorrectShift(t, -10, ciphers.AlphabetLength-10)
}

func Test_DetermineShift_NegativeLarge(t *testing.T) {
	assertCorrectShift(t, -ciphers.AlphabetLength, 0)
	assertCorrectShift(t, -1-ciphers.AlphabetLength, ciphers.AlphabetLength-1)
	assertCorrectShift(t, -1000, 1000%ciphers.AlphabetLength)
}

func assertCorrectShift(t *testing.T, input int, expected int) {
	result := ciphers.DetermineShift(input)

	if result != expected {
		t.Errorf("DetermineShift(%v) == %v, expected %v", input, result, expected)
	}
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
