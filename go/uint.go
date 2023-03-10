package main
// 在 range 里面，数组会拷贝一份，切片则是引用
import (
	"fmt"
	"reflect"
)

func main() {
	var a uint = 0
	var b uint = 1
	c := a - b
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)

	f1()
}

func f1() {
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		fmt.Println("a[k] = ", a[k], "v = ", v)
		a[k] = v + 100
	}
	fmt.Println(a)
}