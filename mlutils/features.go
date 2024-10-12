package mlutils

import (
	"GoNumerica/arei"
	"errors"
)

// Package to handle feature transforms, primarlly using areis

// LabelEncode takes mx1 Arei then returns an mxn Arei with each entry converted into a one-hot label
func LabelEncode(y *arei.Arei) (*arei.Arei, error) {
	// Set up to only process one feature column
	if y.Shape[1] != 1 {
		return nil, errors.New("label encoder only proccess 1d feature columns")
	}

}
