package utils

import "os"

func CheckIfFileExists(fileName string) bool {

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}

	return true
}
