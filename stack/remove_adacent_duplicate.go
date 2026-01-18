package stack

import "fmt"

func removeDuplicates(s string) string {
	stack := []rune{}
	for _, val := range s {
		if len(stack) == 0 {
			stack = append(stack, val)
			continue
		}
		if stack[len(stack)-1] != val {
			stack = append(stack, val)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	res := ""
	for _, ch := range stack {
		res = res + string(ch)
	}
	return res
}

func Duplicate() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("azxxzy"))
}
