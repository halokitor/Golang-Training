package function

import (
	"fmt"
	"math"
)

/* function for two num add */
func addNum_(x, y int) int {
	return x + y
}

/* function for two value swap */
func swapChar_(x, y string) (string, string) {
	return y, x
}

/* function for return struct */
func split_(param float32) (x, y float32) {
	x = param*3.0 + 3.0
	y = float32(math.Sqrt(float64(param)))
	return
}

func PrintFunc() {

	fmt.Println("------------------------------------04_Function------------------------------------")

	/* terminal print add result */
	fmt.Println("add num is", addNum_(9, 8))

	charA := "charA"
	charB := "charB"
	charA, charB = charB, charA

	a, b := swapChar_("charA", "charB")
	fmt.Println("swap char reuslt", a, b, charA, charB)

	fmt.Println(split_(8))

	fmt.Println("-----------------------------------------------------------------------------------")

}
