package main

import "fmt"

type HybridScheduler struct {
    state int
}

func (s *HybridScheduler) run_context(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*54) % 997
    }
    return value
}

func main() {
    obj := &HybridScheduler{state: 54}
    fmt.Println(obj.run_context(54))
}
