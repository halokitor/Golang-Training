package pointerstruct

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

/* "pointer" ctl */
func pointer_(i, j int) {
	p := &i
	fmt.Printf("pointer val: %v\n", *p)

	*p = 21
	fmt.Printf("pointer change val: %v\n", i)

	p = &j
	*p = *p / 37
	fmt.Printf("pointer ctl val: %v\n", j)
}

/* "struct" define and ctl */
type Class struct {
	name   string
	number int
}

/* "pointer" struct define and config */
type PointerStruct struct {
	name string
	num  int
}

/* "array" define and cfg */
func arrayFunc_(array []string) {
	array[0] = "Hello"
	array[1] = "Go!"

	fmt.Println("Array: ", array)
	fmt.Println("Array len: ", len(array))

	/* "slices" Array */
	slices := array[1:3]
	fmt.Printf("Slices Arr: %v,%v, Len: %v\n", slices[0], slices[1], len(slices))

	slicesOne := array[:]
	slicesTwo := array[0:]
	slicesThree := array[:len(array)]
	fmt.Println("slices", slicesOne, slicesTwo, slicesThree)

	/* make "len" and "cap" */
	s := []int{}
	printlnLenCap_(s)

	sInt := []int{1, 3, 45, 6, 3}
	printlnLenCap_(sInt)

	sInt = sInt[:3]
	printlnLenCap_(sInt)

	sInt = sInt[1:]
	printlnLenCap_(sInt)

	makeIntA := make([]int, 6)
	printlnLenCap_(makeIntA)

	makeIntB := make([]int, 0, 6)
	printlnLenCap_(makeIntB)

	/* "slice append" */
	appendInt := []int{}
	appendInt = append(appendInt, 1)
	printlnLenCap_(appendInt)

	appendInt = append(appendInt, 2, 3, 4, 5)
	printlnLenCap_(appendInt)

	/* "Range" test */
	rangeArr := []int{1, 2, 35, 5, 67, 84, 245}
	for index, val := range rangeArr {
		fmt.Printf("index: %d, val: %v\n", index, val)
	}

	powArr := make([]int, 10)
	for i := range powArr {
		powArr[i] = 1 << uint(i)
	}

	for _, val := range powArr {
		fmt.Printf("%d,", val)
	}

}

/* printf len and cap */
func printlnLenCap_(s []int) {
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}

/* "array struct" define and cfg */
type ArrayStruct []struct {
	i int
	b bool
}

/* "slice of slice" define and cfg */
func sliceOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][1] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println("board len: ", len(board))

	for i := 0; i < len(board); i++ {
		fmt.Println("board", board[i])
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
}

/* "Maps" section function */
type MyVertex struct {
	Lat, Long float64
}

var m map[string]MyVertex

func PrintMaps_() {
	m = make(map[string]MyVertex)
	m["Bell Labs"] = MyVertex{
		40.68433, -74.39967,
	}

	var n = map[string]MyVertex{
		"Bell Labs": MyVertex{
			40.68433, -74.39967,
		},
		"Google": MyVertex{
			37.42202, -122.08408,
		},
	}

	var l = map[string]MyVertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	fmt.Println("Bell Labs'Lat & Long", m["Bell Labs"])
	fmt.Println(n, n["Google"])
	fmt.Println(l)

	// judge element if in map
	mm := make(map[string]int)
	mm["request"] = 31
	mm["answer"] = 32

	fmt.Println(mm)
	mm["answer"] = 33
	fmt.Println(mm)
	delete(mm, "answer")
	fmt.Println(mm)

	elem, ok := mm["answer"]
	fmt.Println("Element val:", elem, "Present?", ok)

}

/* "Func Value" trans func */
func Compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func PrintCompute_() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(Compute(hypot))
	fmt.Println(Compute(math.Pow))
}

/* "function-Closures" func */
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func PrintAdder_() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-1*i))
	}
}

/* "exercise-slices" func */
func slicesPic_(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8((x + y) / 2)
		}
		ss[y] = s
	}
	return ss
}

/* "exercise-Maps" function */
func WordCount(s string) map[string]int {
	wordsField := strings.Fields(s)
	m := make(map[string]int)
	count := 0
	for i := 0; i < len(wordsField); i++ {
		for _, s := range wordsField {
			if s == wordsField[i] {
				count++
			}
		}
		m[wordsField[i]] = count
		count = 0
	}
	return m
}

/* "exercies-fibonacci-closure" func */
func fibonacci() func() int {
	prev1, prev2 := 0, 1
	return func() int {
		result := prev1
		prev1, prev2 = prev2, prev1+prev2
		return result
	}
}

func PrintPointerStruct() {

	fmt.Println("---------------------------------06_PointerStruct----------------------------------")

	/* Println : 1 */
	pointer_(42, 2701)

	/* Println : 2 */
	fmt.Println("struct: ", Class{"higher", 40})

	// quest struct
	testStruct := Class{"test", 30}
	fmt.Println("struct.name: ", testStruct.name)

	// change struct
	testStruct.number = 32
	fmt.Println("struct.number: ", testStruct.number)

	/* Println : 3 */
	v := PointerStruct{"pointer", 6}
	p := &v
	/* miss "*", real: *p */
	p.name = "change"
	p.num = 3
	fmt.Println("pointerStruct: ", *p, p.name, v)

	/* Println : 4 */
	inputArr := []string{"o", "k", "!", "lang"}
	arrayFunc_(inputArr)

	/* Println : 5 */
	arrStr := ArrayStruct{{2, true}, {3, false}}
	fmt.Println("ArrayStruct: ", arrStr)

	/* Println : 6 */
	sliceOfSlice()

	/* Println : 7 */
	fmt.Println("slicesPic_: ", slicesPic_(5, 5))
	pic.Show(slicesPic_)

	/* Println : 8 */
	PrintMaps_()

	/* Println : 9 */
	wc.Test(WordCount)

	/* Println : 10 */
	PrintCompute_()

	/* Println : 11 */
	PrintAdder_()

	/* Println : 12 */
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println("-----------------------------------------------------------------------------------")

}
