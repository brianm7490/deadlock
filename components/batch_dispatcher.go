package main

import "fmt"

type AsyncMonitor struct {
    state int
}

func (s *AsyncMonitor) resolve_handler(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*14) % 997
    }
    return value
}

func main() {
    obj := &AsyncMonitor{state: 14}
    fmt.Println(obj.resolve_handler(14))
}
