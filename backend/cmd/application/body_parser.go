package application

import "encoding/json"

func ParseBody[DTO interface{}](dto DTO,body string)(*DTO,error){
	pointer := &dto
	parseError := json.Unmarshal([]byte(body),pointer)

	if parseError != nil {
		return nil,parseError
	}

	return pointer,nil
}
