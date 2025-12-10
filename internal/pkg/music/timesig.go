package music

import "fmt"

var (
	ValidTimeSignatureNumerators   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	ValidTimeSignatureDenominators = []int{2, 4, 8, 16}
)

func nIndexToValue(nIndex int) int {
	for i, n := range ValidTimeSignatureNumerators {
		if i == nIndex {
			return n
		}
	}
	return -1
}

func dIndexToValue(nIndex int) int {
	for i, d := range ValidTimeSignatureDenominators {
		if i == nIndex {
			return d
		}
	}
	return -1
}

type TimeSignature struct {
	Numerator   int
	Denominator int
	nIndex      int // numerator index
	dIndex      int // denominator index
}

func NewTimeSignature() *TimeSignature {
	return &TimeSignature{
		Numerator:   4,
		Denominator: 4,
		nIndex:      nIndexToValue(4),
		dIndex:      dIndexToValue(4),
	}
}

func (t *TimeSignature) String() string {
	return fmt.Sprintf("%d/%d", t.Numerator, t.Denominator)
}

func (t *TimeSignature) IncrementNumerator(dir, count int) {
	tempIndex := t.nIndex + (dir * count)
	if tempIndex >= len(ValidTimeSignatureNumerators) {
		tempIndex = len(ValidTimeSignatureNumerators) - 1
	} else if tempIndex < 0 {
		tempIndex = 0
	}
	t.nIndex = tempIndex
	t.Numerator = ValidTimeSignatureNumerators[tempIndex]
}

func (t *TimeSignature) IncrementDenominator(dir, count int) {
	tempIndex := t.dIndex + (dir * count)
	if tempIndex >= len(ValidTimeSignatureDenominators) {
		tempIndex = len(ValidTimeSignatureDenominators) - 1
	} else if tempIndex < 0 {
		tempIndex = 0
	}
	t.dIndex = tempIndex
	t.Denominator = ValidTimeSignatureDenominators[tempIndex]
}
