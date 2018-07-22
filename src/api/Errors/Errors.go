package Errors

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
