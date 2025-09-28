package main

import (
	"fmt"
	"myFirstGoProject/pacote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(pacote.Foo)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1022}

	for i, elem := range arr {
		fmt.Println(&i, &elem)
	}

}
