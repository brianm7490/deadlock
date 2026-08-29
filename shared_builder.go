package main

import "fmt"

type SecureCollector struct {
    state int
}

func (s *SecureCollector) decode_engine(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*29) % 997
    }
    return total
}

func main() {
    obj := &SecureCollector{state: 29}
    fmt.Println(obj.decode_engine(29))
}
