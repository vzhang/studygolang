package main

import "fmt"

func main() {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	fmt.Println(pc)
}
