/*
base define type in go

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名

	// 表示一个 Unicode 码点

float32 float64

complex64 complex128
*/
package main

import "fmt"

func main() {
	/* Print : 1 */
	// define new single var
	var name string = "kitor"
	var age int = 12
	name = "Alokitor"

	/* Print : 2 */
	// define multi row var
	var (
		firtstName string = "halokitor"
		addr       int    = 6
		email      string = "email.com"
	)

	/* Print : 3 */
	// no use type define, short define can't use out of function
	autoName := "autoName"
	autoNumber := 8

	/* Print : 4 */
	// &: get addr from variable
	autoAddr := 18

	/* Print : 5 */
	// exchange var value

	a := 20
	b := 30

	a, b = b, a
	numA, numB := 10, 20

	/* Print : 6 */
	// var package or func
	var c, python, java bool
	var one, two, three = true, false, "three"

	/* Print : 7 */
	// define const
	const Pi = 3.14159

	fmt.Println("Single Var:", name, age)

	fmt.Println("Multi Var", firtstName, addr, email)

	fmt.Println("Auto Define Var", autoName, autoNumber)

	fmt.Println("Addr: %f, %d", autoAddr, &autoAddr)

	fmt.Println("Exchange Var: %d, %d", a, b)
	fmt.Println("multi define Var: %d, %d", numA, numB)

	fmt.Println("var define package", c, python, java)
	fmt.Println("var define package", one, two, three)

	fmt.Printf("const: type is %T, value is %v", Pi, Pi)
}
