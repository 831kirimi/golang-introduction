package main

import (
	"fmt"
)

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	if IsEven(n) {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

}
