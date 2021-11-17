package initprogram

// applyBitmask2 gets floatingMemLocation and the []int array of all X positions. Obtains all possible X combinations,
// and then it obtains the uint64 representation of all possible memory locations.
func applyBitmask2(memLocation *uint64, bitmask string) []uint64 {

	floatingMemLocation, xPositions := byteArrayWithBitmask2(memLocation, bitmask)

	xPermutations := getBitPermutations(len(xPositions))

	memoryCombinations := getAllMemoryCombinations(&floatingMemLocation, &xPositions, &xPermutations)

	return memoryCombinations
}

// getAllMemoryCombinations loops through all X permutations and applies it one by one to the floatingMemLocation from which it
// obtains uint64 representation, and it attaches it to the []uint64 array that will be returned after all.
func getAllMemoryCombinations(floatingMemLocation *[]byte, xPositions *[]int, xPermutations *[][]byte) []uint64 {

	memCombinations := make([]uint64, 0)

	for _, combination := range *xPermutations {
		for i, x := range combination {
			(*floatingMemLocation)[(*xPositions)[i]] = x
		}
		convertedMemLocation := byteArrToUint64(*floatingMemLocation)
		memCombinations = append(memCombinations, convertedMemLocation)
	}
	return memCombinations
}

// orX
//If the bitmask bit is 0, the corresponding memory address bit is unchanged.
//If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
//If the bitmask bit is X, the corresponding memory address bit is floating.
func orX(memValue byte, mask byte) byte {
	if mask == 'X' {
		return 'X'
	} else {
		return mask | memValue
	}
}

// byteArrayWithBitmask2 converts value to byte array and applies currBitmask with orX operand.
// It also returns the []int array, where X are located in bitmask (==or resultant byte array).
func byteArrayWithBitmask2(value *uint64, bitmask string) ([]byte, []int) {
	bitsValue := integerToBits(*value)
	var byteArr = make([]byte, len(bitmask))
	xPositions := make([]int, 0)

	for i, j := len(bitsValue)-1, len(byteArr)-1; j >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			byteArr[j] = orX(bitsValue[i], bitmask[j])
		} else {
			byteArr[j] = orX(0, bitmask[j])
		}
		if byteArr[j] == 'X' {
			xPositions = append(xPositions, j)
		}
	}

	return byteArr, xPositions
}

// getBitPermutations returns the [][]byte - all n-length permutations of '0' and '1',
// where []byte is the n-length permutation of 0 and 1
func getBitPermutations(n int) [][]byte {
	var result [][]byte
	var tmp = make([]byte, n)

	bitPermutation(n, tmp, &result, n)
	return result
}

// bitPermutation will recursively create permutation of '0' and '1' that are depth variable long and attaches it to
// the result *[][]byte variable.
func bitPermutation(depth int, tmpResult []byte, result *[][]byte, index int) {
	if index == 0 {
		a := make([]byte, depth)
		copy(a, tmpResult)
		*result = append(*result, a)
	} else {
		tmpResult[depth-index] = '0'
		bitPermutation(depth, tmpResult, result, index-1)
		tmpResult[depth-index] = '1'
		bitPermutation(depth, tmpResult, result, index-1)
	}
}
