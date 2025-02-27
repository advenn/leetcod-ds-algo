package main

import (
	"fmt"
	"go_sub/subs/queue_stack"
)

func main() {
	result := queue_stack.IsValid("()[]{}")
	fmt.Printf("result: %v", result)
}
