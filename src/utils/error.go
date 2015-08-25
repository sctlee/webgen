package utils

import "log"

func Check(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
