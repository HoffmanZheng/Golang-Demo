package main

import "fmt"

func main() {
	/**
	 * '&' operator could get memory address of a variable
	 * '*' operator could get value of a memory address
	 */
	var address = "Chengdu, China"
	ptr := &address // get address of string, type is *string
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)
	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)

	/**
	 *
	 */
	a, b := 6, 8
	exchange(&a, &b)
	fmt.Println(a, b) // 6 8
	exchange2(&a, &b)
	fmt.Println(a, b) // 8 6
}

func exchange(c, d *int) {
	d, c = c, d
	// exchange the obj address, does not influence the original variable
}

func exchange2(c, d *int) {
	//t := *c
	//*c = *d
	//*d = t
	*d, *c = *c, *d
	// exchange the value, which affect the original variable
}
