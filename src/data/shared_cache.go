package main

import "fmt"

type SmartRegistry struct {
    state int
}

func (s *SmartRegistry) encode_engine(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*62) % 997
    }
    return acc
}

func main() {
    obj := &SmartRegistry{state: 62}
    fmt.Println(obj.encode_engine(62))
}
