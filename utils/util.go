package utils

import "log"

func CheckError(e error, msg ...interface{}) {
	if e != nil {
		log.Fatal(msg)
		return
	}
}
