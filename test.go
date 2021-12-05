package main

import "fmt"

var variable int8 = 20

func test() {
	//if variable > 30 {
	//	fmt.Printf("%d gt: %d", variable, 30)
	//} else if variable > 40 {
	//	fmt.Printf("%d gt: %d", variable, 40)
	//} else {
	//	fmt.Printf("other value: %d", variable)
	//}

	switch {
	case variable > 30:
		fmt.Printf("%d gt: %d", variable, 30)
	case variable > 40:
		fmt.Printf("%d gt: %d", variable, 40)
	default:
		fmt.Printf("other value: %d", variable)
	}
	//switch variable {
	//case 20:
	//	print("bing")
	//case 30:
	//	print("seen")
	//case 40:
	//	print("teen")
	//default:
	//	print("pink")
	//}

}

func add(a int, b int) int {
	return a + b
}

func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

//func main() {
//for i := 0; i < 50; i++ {
//	println(i,fib(i))
//}
//fmt.Printf("for what?")
//test()
//print(test())
//var strArr = []int{1, 2, 3}
//println(add(1, 2))
//for i, item := range strArr {
//	fmt.Println(i, item)
//}

//}
