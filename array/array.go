package main

import "fmt"

func List() [5]int {
	return [5]int{1, 2, 3, 4, 5}
}

func main() {
	list := List()
	fmt.Printf("%v\n", list)
}
