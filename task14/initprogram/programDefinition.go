package initprogram

// Instructions example:
// mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X string of 36 bits, written with the most significant bit (representing 2^35) on the left
// mem[8] = 11 Values and memory addresses are both 36-bit unsigned integers.
type Instructions struct {
	currBitmask  string
	memLocations map[uint64]uint64
}

func InitializeInitProgram() *Instructions {
	i := &Instructions{}
	i.memLocations = make(map[uint64]uint64, 0)

	return i
}

func (i *Instructions) MemLocations() map[uint64]uint64 {
	return i.memLocations
}

func (i *Instructions) UpdateBitmask(newBitmask string) {
	i.currBitmask = newBitmask
}

// Write Subtask 1 -> bitmask is applied on value.
func (i *Instructions) Write(memLocation uint64, value uint64) {
	applyBitmask(&value, i.currBitmask)
	i.memLocations[memLocation] = value
}

func (i *Instructions) SumOfMemLocations() uint64 {
	result := uint64(0)
	for _, value := range i.memLocations {
		result += value
	}
	return result
}

// WriteSubtask2 Subtask 2 -> bitmask is applied on memLocation.
func (i *Instructions) WriteSubtask2(memLocation uint64, value uint64) {
	memoryCombinations := applyBitmask2(&memLocation, i.currBitmask)

	for _, memLoc := range memoryCombinations {
		i.memLocations[memLoc] = value
	}
}
