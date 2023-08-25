package process

import (
	"fmt"
	"math"
	"time"
)

/* "for" process function */
func sum_(x int) int {
	for i := 0; i < 10; i++ {
		x += i
	}
	return x
}

/* "short for" process func */
func shortSum_(x, y int) int {
	for y < 50 {
		x += 1
		y += 1
	}
	return x
}

/* "if" process func */
func if_(x, y, lim float64) (result float64) {
	if val := math.Pow(x, y); val < lim {
		result = val
		return
	}
	result = 2
	return
}

/* "if else" process func */
func ifelse_(x, y, lim float64) {
	if val := math.Pow(x, y); val < lim {
		fmt.Printf("%g < %g\n", val, lim)
	} else if val == lim {
		fmt.Printf("%g == %g\n", val, lim)
	} else {
		fmt.Printf("%g > %g\n", val, lim)
	}
}

/* circle problem func */
func newTonCircle(x, z float64) (float64, float64) {
	res := math.Sqrt(x)
	count := 0.0

	for math.Abs(float64(z-res)) > 0.01 {
		z -= (z*z - x) / (2 * z)
		count += 1
		if count >= 500 {
			return z, count
		}
	}

	return z, count
}

/* "switch" process func */
func switchProcess() {
	defer fmt.Println("world one")
	defer fmt.Println("world two")
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func PrintProcess() {

	fmt.Println("-------------------------------------05_Process------------------------------------")

	/* Printf : 1 */
	fmt.Printf("sum_type: %T, sum_val: %v\n", sum_(1), sum_(1))

	/* Printf : 2 */
	fmt.Println("sumShort val: ", shortSum_(10, 20))

	/* Prinln : 3 */
	fmt.Println("if val: ", if_(2, 3, 9), if_(2, 3, 5))

	/* Prinln : 4 */
	ifelse_(2, 3, 9)
	ifelse_(2, 3, 5)
	ifelse_(2, 3, 8)

	/* Prinln : 5 */
	a, b := newTonCircle(3, 1)
	fmt.Printf("x = %v, z = 1, z= %v, count = %v\n", math.Sqrt(3), a, b)

	/* Prinln : 6 */
	switchProcess()

	fmt.Println("-----------------------------------------------------------------------------------")

}
