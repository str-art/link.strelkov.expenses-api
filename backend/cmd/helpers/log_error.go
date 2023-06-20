package helpers

import "log"

func LogError(err error){
	log.Printf("Recieved error: %#v",err)
}