package mypkg

import (
	"fmt"
	"math"
	"math/rand"
)

func PrintPkg() {

	fmt.Println("----------------------------------02_ImportPkg-------------------------------------")

	fmt.Println("Import Package, random num is", rand.Intn(3))
	fmt.Println("Import Package, sqrt num is", math.Sqrt(9))
	fmt.Println("Import Package, math Pi is", math.Pi)

	fmt.Println("-----------------------------------------------------------------------------------")

}
