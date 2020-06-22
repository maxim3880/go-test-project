package acronym

import (
	"fmt"
	"strings"
)

//Execute start of presentation task 1
func Execute() {
	fmt.Println(getAcronymText("portable network graphics "))
}

func getAcronymText(inputString string) (result string) {
	arr := strings.Fields(inputString)
	for _, value := range arr {
		if value != "" {
			result += strings.ToUpper(string(value[0]))
		}
	}
	return
}
