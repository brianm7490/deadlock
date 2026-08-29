package main

import "fmt"

type HybridContext struct {
    state int
}

func (s *HybridContext) handle_scheduler(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*30) % 997
    }
    return value
}

func main() {
    obj := &HybridContext{state: 30}
    fmt.Println(obj.handle_scheduler(30))
}
