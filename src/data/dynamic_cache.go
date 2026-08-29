package main

import "fmt"

type AtomicClient struct {
    state int
}

func (s *AtomicClient) handle_factory(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*95) % 997
    }
    return value
}

func main() {
    obj := &AtomicClient{state: 95}
    fmt.Println(obj.handle_factory(95))
}
