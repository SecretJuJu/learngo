package basic

import (
	"fmt"
	"github.com/secretjuju/learngo/something"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

// naked return
func nakedReturn(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer
func deferExample() {
	defer fmt.Println("I'm done1")
	fmt.Println("Start")
	defer fmt.Println("I'm done2")
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// koreanAge 가 if 문 안에서만 사용한다는 것을 표현할 수 있다.
	if koreanAge := age + 2; koreanAge < 19 {
		return false
	}

	// fmt.Println(koreanAge) // koreanAge is not available here
	return true
}

func RunBasics() {
	fmt.Println("----- Basic -----")

	fmt.Println("Hello, world!")
	something.SayHello()

	const name string = "secretjuju"
	const age int = 22
	const isAdult bool = true

	fmt.Println(name, age, isAdult)

	var description = "I am a student"
	fmt.Println(description)

	description = "I am a developer"
	fmt.Println(description)

	gender := "male"
	fmt.Println(gender)

	fmt.Printf("multiply: %d\n", multiply(2, 2))

	const sourceName = "secretjuju"
	length, uppercase := nakedReturn(sourceName)
	fmt.Println(length, uppercase)

	// after strings upper, original string is not changed
	fmt.Println(sourceName) // secretjuju

	// defer
	// defer 은 Stack 처럼 동작하는듯 하다.
	deferExample() // Start, I'm done2, I'm done1

	// [1,2,3,4,5]
	addSources := []int{1, 2, 3, 4, 5}
	addResult := superAdd(addSources...)
	// [1,2,3,4,5] => 15
	fmt.Println(addSources, "=>", addResult)

	// canIDrink
	fmt.Println(canIDrink(16)) // false
	fmt.Println(canIDrink(18)) // true
}
