package initprogram

import (
	"math"
	"strconv"
)

// applyBitmask obtains the value on which bitmask is applied and assigns it to the existing value.
func applyBitmask(value *uint64, bitmask string) {
	// 32 bit long array [32]byte
	byteValue := byteArrayWithBitmask(value, bitmask)
	*value = byteArrToUint64(byteValue)
}

// byteArrToUint64 converts byte array to uint64.
func byteArrToUint64(value []byte) uint64 {
	result := 0.0
	for i, power := len(value)-1, 0.0; i >= 0; i, power = i-1, power+1 {
		if value[i] == 1 || string(value[i]) == "1" {
			result += math.Pow(2, power)
		}
	}
	return uint64(result)
}

// byteArrayWithBitmask converts value to byte array and applies currBitmask with xand operand.
func byteArrayWithBitmask(value *uint64, bitmask string) []byte {
	bitsValue := integerToBits(*value)
	var byteArr = make([]byte, len(bitmask))

	for i, j := len(bitsValue)-1, len(byteArr)-1; j >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			byteArr[j] = xand(bitsValue[i], bitmask[j])
		} else {
			byteArr[j] = xand(0, bitmask[j])
		}
	}
	return byteArr
}

// xand operand:
// Mask is composed of 1 | 0 | X
// X in currBitmask and whatever value is that value.
// 1 | 0 in currBitmask rewrites the value corresponding to currBitmask value.
func xand(val byte, mask byte) byte {
	if mask == 'X' {
		return val
	} else {
		return mask - 48 // byte to bin
	}
}

// integerToBits gives the byte array, where each byte represents 0 | 1.
func integerToBits(n uint64) []byte {

	tmpResult := []byte(strconv.FormatUint(n, 2))

	for i := len(tmpResult) - 1; i >= 0; i = i - 1 {
		tmpResult[i] -= 48
	}
	return tmpResult
}
