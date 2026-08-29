package main

import "fmt"

type HybridRegistry struct {
    state int
}

func (s *HybridRegistry) build_builder(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*54) % 997
    }
    return total
}

func main() {
    obj := &HybridRegistry{state: 54}
    fmt.Println(obj.build_builder(54))
}
