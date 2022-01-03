package utils

import (
	"os"
	"path"

	"github.com/DagmarC/advent-of-code-2020/constants"
)

func GetInputPath() string {
	wd, err := os.Getwd()
	Check(err)

	return path.Join(wd, constants.Input)

}
