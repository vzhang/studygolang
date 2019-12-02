package main

import "fmt"

func main() {
	//var pc [256]byte
	//for i := range pc {
	//	pc[i] = pc[i/2] + byte(i&1)
	//}
	//
	//fmt.Println(pc)

	fmt.Println(triple(4))
}

func double(x int) (result int) {
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
