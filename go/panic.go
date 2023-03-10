// generate a main function
package main

import (
	"fmt"
)

func f() (r string) {
	a := make([]int, 0)
	defer func() {
		fmt.Println("defer0")
			r = "f return error"
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	a[2] = 1
	defer func() {
		fmt.Println("defer1")
	}()

	// defer func() {
	fmt.Println("defer2")
	if r := recover(); r != nil {
		fmt.Println("Recovered in f", r)
	}
	// }()
	fmt.Println("Calling g.")
	fmt.Println("Returned normally from g.")
	r = "f return ok"
	return
}

func main() {
	fmt.Println("Hello World")
	defer func() {
		fmt.Println("defer3")
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("f return ", f())
	defer func() {
		fmt.Println("defer4")
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Byte World")
}
