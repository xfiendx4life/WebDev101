package main


import (
	"errors"
	"fmt"
)

// // https://leetcode.com/problems/add-two-numbers/
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

//  type ListNode struct {
// 	     Val int
// 	     Next *ListNode
//  }

//  func (ln *ListNode) Len() int {
//     for n := 0; ln.Next != nil; n++ {}
//     return n
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     var p int
//     for
// }

// func f(fun func(x int32) int32, a ...int32) []int32 {
// 	res := make([]int32, 0, len(a))
// 	for _, item := range a {
// 		res = append(res, fun(item))
// 	}
// 	return res
// }

type stack struct {
	data []int
}

func (s *stack) push(a int) {
	s.data = append(s.data, a)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) pop() (int, error) {
	if !s.isEmpty() {
		a := s.data[len(s.data)-1]
		s.data = s.data[0 : len(s.data)-1]
		return a, nil
	}
	return 0, errors.New("stack is empty")
}

func (s *stack) back() (int, error) {
	if !s.isEmpty() {
		a := s.data[len(s.data)-1]
		return a, nil
	}
	return 0, errors.New("stack is empty")
}

func (s *stack) length() int {
	return len(s.data)
}

type queue struct {
    stack
}

func second() {
	// a := [6]int32{1, 2, 3, 15, 150, 4}
	// a := []int32{1, 2, 3, 15, 150, 4}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%d ", a[i])
	// }
	// fmt.Println()
	// for i, item := range a {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%d ", item)
	// 	}
	// }
	// fmt.Println()
	// a := make([]int32, 0, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, int32(i*10+i))
	// 	fmt.Printf("%v len = %d, cap = %d\n", a, len(a), cap(a))
	// }
	// b := a[3:5:10]
	// b[1] = 150
	// fmt.Println(a)

	// a = append(a[:3], a[4:]...)
	// fmt.Println(a)

	// fmt.Println(f(func(x int32) int32 {
	// 	var s int32
	// 	for s = 0; x > 0; x /= 10 {
	// 		s += int32(x % 10)
	// 	}
	//     return s
	// }, a...))

	// fmt.Println(f(func(x int32) int32 {
	//     return x % 10
	// }, a...))

	// a := make(map[int]int)
	// r := make([]int, 10)
	// rand.Seed(time.Now().Unix())
	// for i := range r {
	// 	r[i] = rand.Intn(10)
	// }
	// for _, item := range r {
	// 	if _, ok := a[item]; !ok {
	// 		a[item] = 1
	// 	} else {
	// 		a[item]++
	// 	}
	// }
	// for k, v := range a {
	//     fmt.Printf("%d %d\n", k, v)
	// }

	// a := make()

	s := stack{}
	s.push(5)
	res, err := s.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	_, err = s.back()
	if err != nil {
		fmt.Println(err)
	}

    q := queue{}
    q.pop()
}
