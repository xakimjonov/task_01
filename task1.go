package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)

}
