package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("canon")
	l.PushFront(67)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
