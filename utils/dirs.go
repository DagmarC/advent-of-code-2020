package utils

import "os"

func GetWd() string {
	wd, err := os.Getwd()
	Check(err)
	return wd
}
