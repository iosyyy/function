package main

import (
	"fmt"
	_ "function/src/github.com/iosyyy"
	_ "os"
)

func test() {
	a := 666
	fmt.Println("hello func", a)
}
func test1(name string) {
	fmt.Println(name)

}
func test2(name string, types ...int) {
	test3(name, types[:2]...)
	test3(name, types[1:4]...)

}

func test3(name string, types ...int) {
	fmt.Println(name)

	for _, value := range types {
		fmt.Println(value)
	}
	fmt.Println()

}
func test4(name string) (nameClone string) {
	nameClone = name
	return
}

func test5(name string) (nameClone func() string) {
	nameClone = func() string {
		fmt.Println(name)
		name += "hello "
		return name
	}
	return
}

func test6(a int) (sum int) {
	if a == 0 {
		sum = a
		return
	}
	sum = a + test6(a-1)
	defer fmt.Println("end test6")
	return
}

func getMaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

type nameX func(int, int) (int, int)

func main() {
	/*funcX := test5("hello world")
	funcX()
	funcX()
	funcX()
	funcX()
	funcX()*/
	//test6(100)
	//args := os.Args
	//for _, arg := range args {
	//	fmt.Println(arg)
	//}
}
