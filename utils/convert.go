package utils

import "strconv"

func ToInt(s string) int {
	number, err := strconv.Atoi(s)
	Check(err)
	return number
}
