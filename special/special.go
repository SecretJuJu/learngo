package special

import "fmt"

// get pointer, change value
func changedInFunc(a *int) {
	*a = 10
}

func RunSpecial() {
	fmt.Println("----- Special -----")
	// pointer
	a := 2
	aa := &a

	// *aa 는 aa 가 가리키는 주소에 있는 값을 의미한다.
	*aa = 3
	fmt.Println(a)                                 // 3
	fmt.Printf("aa === &a => %p === %p\n", aa, &a) // true
	changedInFunc(&a)
	fmt.Println(a) // 10

	// array
	names := [5]string{"secretjuju1", "secretjuju2", "secretjuju3", "secretjuju4"}
	names[4] = "secretjuju5"
	fmt.Println(names) // [secretjuju1 secretjuju2 secretjuju3 secretjuju4 secretjuju5]

	// append don't change original array (slice)
	// names 를 복제한 끝이 정해지지 않은 배열 names1 을 만든다.
	names1 := append(names[:])
	names2 := append(names1, "secretjuju6")
	fmt.Println(names)  // [secretjuju1 secretjuju2 secretjuju3 secretjuju4 secretjuju5]
	fmt.Println(names2) // [secretjuju1 secretjuju2 secretjuju3 secretjuju4 secretjuju5 secretjuju6]

	// map
	secretjuju := map[string]string{"name": "secretjuju", "age": "22"}
	fmt.Println(secretjuju) // map[age:22 name:secretjuju]

	// map 을 iterate 하기 위해서는 key, value 를 받아서 사용해야 한다.
	for key, value := range secretjuju {
		fmt.Println(key, value)
	}

	type person struct {
		name    string
		age     int
		favFood []string
	}

	favFoods := []string{"taco"}
	secretjuju1 := person{"secretjuju", 22, favFoods}
	fmt.Printf("who is %s? %v\n", secretjuju1.name, secretjuju1)

	// 명시적인 필드 값 지정
	favFoods2 := []string{"kimchijjigae"}
	secretjuju2 := person{
		name:    "secretjuju",
		age:     22,
		favFood: favFoods2,
	}

	fmt.Printf("who is %s? %v\n", secretjuju2.name, secretjuju2)
}
