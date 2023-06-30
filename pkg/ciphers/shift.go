package ciphers

import (
	"fmt"
)

const LowerBoundLC = 'a'
const UpperBoundLC = 'z'
const LowerBoundUC = 'A'
const UpperBoundUC = 'Z'

const AlphabetLength = int(UpperBoundLC - LowerBoundUC)

func ShiftRune(input rune, shift int) (*rune, error) {
	shift %= AlphabetLength

	isASCII := input >= LowerBoundLC && input <= UpperBoundLC || input >= LowerBoundUC && input <= UpperBoundUC

	if !isASCII {
		return nil, fmt.Errorf("not an ascii character: '%c'", input)
	}

	result := input + rune(shift)

	return &result, nil
}
