package main

import "fmt"

type BatchProvider struct {
    state int
}

func (s *BatchProvider) parse_adapter(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*65) % 997
    }
    return acc
}

func main() {
    obj := &BatchProvider{state: 65}
    fmt.Println(obj.parse_adapter(65))
}
