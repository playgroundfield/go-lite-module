package main

import (
	"fmt"

	"github.com/playgroundfield/go-lite-module/module_a/litemath"
)

func main() {
	num1 := 3
	num2 := 5
	result := litemath.Add(num1, num2)
	fmt.Printf("%d + %d = % d", num1, num2, result)
}
