package main

import (
	"fmt"
)

// slice 는 동적 배열 타입니다. 배열과 달리 크기를 미리 정하지 않아도 되며, 요소를 추가하거나 삭제할 수 있습니다다

func main() {
	var slice = []int{1, 2, 3}
	fmt.Println(slice)
	if slice != nil {
		fmt.Println("not nil")
	}
	slice[1] = 10
	fmt.Println(slice)

	// slice 는 배열을 가리키는 포인터이다
	// 따라서 slice 를 다른 변수에 할당하면 같은 배열을 가리키게 된다
	var slice2 = slice
	println(slice2)

	var slice3 = make([]int, 5, 10)
	fmt.Println("Initial slice:", slice3)
	fmt.Println("Initial length:", len(slice3))
	fmt.Println("Initial capacity:", cap(slice3))

	// Adding elements to exceed the initial capacity
	for i := 0; i < 10; i++ {
		slice3 = append(slice3, i)
		fmt.Printf("After appending %d: length = %d, capacity = %d\n", i, len(slice3), cap(slice3))
	}

	var slice4 = make([]int, 5)
	slice4 = append(slice4, 1, 2, 3, 4)
	fmt.Println(slice4)

	for i, v := range slice4 {
		fmt.Println("index : ", i)
		slice4 = append(slice4, v)
		fmt.Println("v:", v)
	}
	fmt.Println(slice4)

	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	slice5 := []int{1, 2, 3, 4, 5, 6}

	changeArray(&arr1)
	changeSlice(slice5)

	fmt.Println(arr1) // 변하지 않는다
	fmt.Println(slice5)

}

func changeArray(arr *[6]int) {
	arr[2] = 100
}

func changeSlice(slice []int) {
	slice[2] = 100
}

// slice 는 배열을 가리키는 포인터이다 구조체로 구현되어 있으며, 길이와 용량을 가지고 있다 data 필드는 포인터이며, 실제 데이터를 가리킨다
// 따라서 slice 를 다른 변수에 할당하면 같은 배열을 가리키게 된다
