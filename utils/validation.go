package utils

import "strconv"

func IsEmpty(val string) bool {
	if len(val) > 0 {
		return false
	}
	return true
}

func IsNumber(val string) bool {
	_, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return true
}
