package main

import (
	"fmt"
	"runtime"
)

func main() {
	/* terminal print string */
	fmt.Println("Hello Golang!\n", "The next words!")
	fmt.Println(runtime.Version())
}
