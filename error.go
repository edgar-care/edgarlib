package edgarlib

import "fmt"

func CheckError(err error) {
	fmt.Print("test")
	if err != nil {
		panic(err)
	}
}
