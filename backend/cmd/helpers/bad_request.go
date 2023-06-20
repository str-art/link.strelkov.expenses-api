package helpers

import "fmt"

func BadRequestMessage(message string)(string){
	return fmt.Sprintf("%s: %s","Bad request",message)
}
