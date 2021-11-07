package main

import
(
	"fmt"
	"os"
)

func main() {
	// echo what you type as param after go run 2.echo.go
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)
}
