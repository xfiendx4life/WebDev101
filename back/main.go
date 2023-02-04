package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
	
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println("Hello, World")
	fmt.Printf("Hello, web %T\n", 101)
	var a int
	a = 10
	fmt.Println(a)
	b := 4
	fmt.Printf("%d %T\n", b, b)
	fmt.Println(a + b)
	fmt.Println(a*b, a/b, a%b, float64(a)/float64(b), a-b)

	s := "хello"
	fmt.Printf("len = %d, rune count = %d\n", len(s), utf8.RuneCountInString(s))
	r := 'х'
	fmt.Println(r)
	fmt.Println(s[3])
	for i := 0; i < utf8.RuneCountInString(s); i++ {
		if i == 3 {
			continue
		}
		fmt.Println(s[i])
	}
	a = 123
	sum := 0
	for sum = 0; a > 0; a /= 10 {
		sum += a % 10
	}
	fmt.Println(sum)

	if sum > 10 {
		fmt.Println("sum > 10")
	} else if a < 5 {
		fmt.Println("hello")
	}

	rand.Seed(time.Now().Unix())
	a = rand.Intn(11)
	switch a {
	case 5:
		fmt.Println("In the middle")
	case 10:
		fmt.Println("impossible")
	default:
		fmt.Printf("something usual %d\n", a)
	}

	// _, err  := fmt.Scan(&a)
	a = 5
	// if err != nil {
	// 	fmt.Printf("error accured %s\n", err)
	// }
	// var arr []int
	// fmt.Println(arr[3])

	second()
}
