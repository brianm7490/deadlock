package main

import "fmt"

type HybridRegistry struct {
    state int
}

func (s *HybridRegistry) fetch_provider(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*28) % 997
    }
    return acc
}

func main() {
    obj := &HybridRegistry{state: 28}
    fmt.Println(obj.fetch_provider(28))
}
