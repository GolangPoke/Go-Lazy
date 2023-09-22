package log

import "fmt"

func ThrowError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
