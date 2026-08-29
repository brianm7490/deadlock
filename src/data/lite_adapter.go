package main

import "fmt"

type BatchManager struct {
    state int
}

func (s *BatchManager) sync_controller(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*4) % 997
    }
    return value
}

func main() {
    obj := &BatchManager{state: 4}
    fmt.Println(obj.sync_controller(4))
}
