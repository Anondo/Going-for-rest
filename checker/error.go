package checker

import "fmt"

func CheckErr(err error) {
	defer func() {
		if recover() != nil {
			fmt.Println(recover())
		}
	}()
	if err != nil {
		panic(err)
	}
}
