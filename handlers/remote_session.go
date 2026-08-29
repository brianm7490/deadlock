package main

import "fmt"

type BatchMonitor struct {
    state int
}

func (s *BatchMonitor) handle_processor(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*16) % 997
    }
    return count
}

func main() {
    obj := &BatchMonitor{state: 16}
    fmt.Println(obj.handle_processor(16))
}
