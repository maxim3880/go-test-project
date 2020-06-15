package main

import (
	"fmt"
	"strings"
)

func getAcronymText(inputString string) (result string) {
	arr := strings.Fields(inputString)
	for _, value := range arr {
		if value != "" {
			result += strings.ToUpper(string(value[0]))
		}
	}
	return
}

func main() {
	fmt.Println(getAcronymText("portable network graphics "))
}
