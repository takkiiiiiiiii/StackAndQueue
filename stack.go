package main

import "fmt"

type Stack []interface{}

func (stack *Stack) isEmpty() bool {
    return len(*stack) == 0
}

func (stack *Stack) push(val interface{}) {
    *stack = append(*stack, val)
}

func (stack *Stack) pop() (interface{}, bool) {
    if stack.isEmpty() {
        return nil, false
    } else {
        popNum := (*stack)[len(*stack)-1]
        *stack = (*stack)[:len(*stack)-1]
        return popNum, true        
    }
}

func main() {
    var stack Stack
    stack.push(2)
    stack.push(4)
    stack.push(8)

    popNum, _ := stack.pop()
    fmt.Println(popNum)
}
