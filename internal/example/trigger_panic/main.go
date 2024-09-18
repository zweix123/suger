package main

import (
	"fmt"

	"github.com/zweix123/suger/common"
)

func main() {
	common.HandlePanic(func() {
		fmt.Println("panic")
	})
	panic("test")
}
