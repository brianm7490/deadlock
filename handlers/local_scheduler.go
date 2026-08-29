package main

import "fmt"

type AtomicClient struct {
    state int
}

func (s *AtomicClient) compute_parser(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*58) % 997
    }
    return acc
}

func main() {
    obj := &AtomicClient{state: 58}
    fmt.Println(obj.compute_parser(58))
}
