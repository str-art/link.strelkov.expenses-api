package helpers

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)


func StringToCCase(s string)(string){
	regexp := regexp.MustCompile(`\s+`)
	
	trimmedString := regexp.ReplaceAllString(s,"_")

	return strings.ToUpper(trimmedString)
}

func CCaseToString(s string)(string){
	stringWithoutUnderscores := strings.ReplaceAll(s,"_"," ")
	return cases.Title(language.AmericanEnglish,cases.Compact).String(stringWithoutUnderscores)
}