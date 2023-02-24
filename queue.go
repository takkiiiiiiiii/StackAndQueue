package main

import (
    "fmt"
)

type Queue []interface{}

func (queue *Queue) isEmpty() bool {
    return len(*queue) == 0
}

func (queue *Queue) enqueue(val interface{}) {
    *queue = append(*queue, val)
}

func (queue *Queue) dequeue() (interface{}, bool){
    if queue.isEmpty() {
        return nil, false
    } else {
        popNum := (*queue)[0]
        *queue = (*queue)[1:]
        return popNum, true
    }
}

func main() {
    var queue Queue
    queue.enqueue(1)
    queue.enqueue(2)
    queue.enqueue(3)
    
    popNum, _ := queue.dequeue()
    fmt.Println(popNum)
    fmt.Println(queue)
}
