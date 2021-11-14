package main

import "fmt"

// 切片的动态扩容
func main() {
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}

	sli1 := arr1[0:2] // len=2,cap=4
	sli2 := arr2[2:4] // len=2,cap=2

	fmt.Printf("sli1 pointer is [%p], len is [%v], cap is [%v], value is %v \n", &sli1, len(sli1), cap(sli1), sli1)
	fmt.Printf("sli2 pointer is [%p], len is [%v], cap is [%v], value is %v \n", &sli2, len(sli2), cap(sli2), sli2)

	newSli1 := append(sli1, 5)
	newSli2 := append(sli2, 5)

	fmt.Printf("newSli1 pointer is [%p], len is [%v], cap is [%v], value is %v \n",
		&newSli1, len(newSli1), cap(newSli1), newSli1)
	fmt.Printf("source arr1 become: %v \n", arr1)
	fmt.Printf("newSli2 pointer is [%p], len is [%v], cap is [%v], value is %v \n",
		&newSli2, len(newSli2), cap(newSli2), newSli2)
	fmt.Printf("source arr2 become: %v \n", arr2)
}
