package helpers

import "fmt"

func HandleError(str string, err error) {
	if err != nil {
		fmt.Println(str+":", err)
	}
}
