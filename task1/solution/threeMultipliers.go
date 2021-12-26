package solution

import (
	"errors"
	"fmt"

	"github.com/DagmarC/advent-of-code-2020/datafile"
)

func Fin3dMultipliers(sum int) (int, error) {
	var data []int
	datafile.LoadFileIntoSlice(&data)

	tempSum := 0
	var x, y, z int

outer:
	for i := 0; i < len(data); i++ {
		tempSum = data[i]

		for j := i + 1; j < len(data); j++ {
			tempSum = data[i] + data[j]
			if tempSum >= sum {
				continue
			}

			for k := j + 1; k < len(data); k++ {
				tempSum = data[i] + data[j] + data[k]

				if tempSum == sum {
					x, y, z = data[i], data[j], data[k]
					break outer

				} else if tempSum > sum {
					continue
				}
			}
		}
	}
	if x == 0 && y == 0 && z == 0 {
		return 0, errors.New("no match found")
	}
	fmt.Println("Resultant numbers: ", x, y, z)
	return x * y * z, nil
}
