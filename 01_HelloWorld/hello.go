package hello

import (
	"fmt"
	"runtime"
)

func Hello_(name string) {

	fmt.Println("---------------------------------01_HelloWorld-------------------------------------")

	/* terminal print string */
	fmt.Printf("Hello %v, Welcome to Golang!\nThe next words!:)\n", name)
	fmt.Println(runtime.Version())

	fmt.Println("-----------------------------------------------------------------------------------")

}
