package main

import "fmt"

func main() {

	fmt.Println("basic programs 📚")

	/* // p1 add two numbers

	var a, b int

	fmt.Println("enter two numbers")
	fmt.Scanln(&a, &b)

	fmt.Println("sum = ", a+b)
	*/

	// **********************************************************************************

	/* // p2 conditionals

	num1 := 10
	num2 := 20

	if num1 > num2 {
		fmt.Println("num1 is greater")
		} else {
			fmt.Println("num2 is greater")
			} */

	// **********************************************************************************

	/* // p3 loops (for,while)

	for i := 0; i < 5; i++ {
		fmt.Print(i, "\t")
	}

	fmt.Println()

	x := 5

	for x > 0 { //while loop in go declared as this
		fmt.Print(x, "\t")
		x--

	} */

}

// basics misc

/*
	// variable , const , type conversion

	func main() {

	var num_1 int64 = 24
	fmt.Println(num_1)

	num_2, num_3 := 10, 11
	fmt.Println(num_2, num_3)

	var myInt int
	var myString string
	var myFloat float32
	var myBool bool

	fmt.Println(myInt, myString, myFloat, myBool)

	const pi float32 = 3.1412
	fmt.Println(pi)

	a, b, c := 12, 2.5, 2

	fmt.Println(a / int(b))
	fmt.Println(b * float64(c))

}


*/

/*

	// Functions , Error Handling , Switch case

	func main() {

	var quotient, remainder, err = divideNumber(10, 6)

	switch {

	case err != nil:
		fmt.Printf(err.Error() + "\n")

	case remainder == 0:
		fmt.Printf("quotient is %v \n", quotient)

	default:
		fmt.Printf("quotient is %v and remainder is %v \n", quotient, remainder)

	}

	switch remainder {
	case 0:
		fmt.Println("division was successful")

	case 1, 2:
		fmt.Println("division was close")

	default:
		fmt.Println("division was not close")

	}

}

func printMe(inputString string) {
	fmt.Println(inputString)
}

func divideNumber(numerator int, denomainator int) (int, int, error) {
	var err error

	if denomainator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	result := numerator / denomainator
	remainder := numerator % denomainator

	return result, remainder, err
}
*/

/*

	// Arrays , Slices , Maps , loops control structures

		func main() {

	var arr1 [5]int
	fmt.Println(arr1)

	arr2 := [5]int{1, 3, 5, 7, 11}
	fmt.Println(arr2)

	arr3 := [...]int{1, 3, 5, 7, 11, 13, 17, 23}
	fmt.Println(arr3)

	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i], " ")
	}

	fmt.Println()

	for index, value := range arr3 {
		fmt.Print(index, "->", value, " ")
	}

	fmt.Println()

	for _, value := range arr1 {
		fmt.Print(value, " ")
	}

	fmt.Println()

	intSlice := []int32{4, 8, 12}

	fmt.Printf("length of slice is %v and capacity is %v \n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)

	fmt.Printf("length of slice is %v and capacity is %v \n", len(intSlice), cap(intSlice))

	fmt.Println(intSlice)

	var myMap map[string]uint16 = make(map[string]uint16)

	fmt.Println(myMap)

	var myMap2 = map[string]uint16{"Adam": 23, "Sarah": 32, "Harshit": 22}

	fmt.Println(myMap2)

	var age, ok = myMap2["Harshit"]

	if ok {
		fmt.Println("Harshit age is", age)
	} else {
		fmt.Println("Harshit not found in map")
	}

	for name, age := range myMap2 {
		println(name, age)
	}

}

*/

/*

	// Strings , Runes , Bytes

*/

/*

// Structs and Interfaces

type gasEngine struct {
	mpg     uint8
	gallons uint8
}


func (e *gasEngine) setGasEgnine(mpg uint8, gallons uint8) {
	e.mpg = mpg
	e.gallons = gallons
}

func (e gasEngine) getMilesLeft() (miles uint16) {

	return uint16(e.gallons * e.mpg)

}

func (e gasEngine) canMakeIt(milesToCover uint8) bool {
	if e.getMilesLeft() < uint16(milesToCover) {
		return false
	}

	return true

}


func main() {

	var myGasEngine gasEngine

	myGasEngine.setGasEgnine(10, 25)

	fmt.Printf("%+v \n", myGasEngine)

	fmt.Println(myGasEngine.getMilesLeft())

	fmt.Println(myGasEngine.canMakeIt(251))

}


*/
