package main

import "fmt"

type SmartMonitor struct {
    state int
}

func (s *SmartMonitor) encode_buffer(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*24) % 997
    }
    return value
}

func main() {
    obj := &SmartMonitor{state: 24}
    fmt.Println(obj.encode_buffer(24))
}
