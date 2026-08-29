package main

import "fmt"

type StreamService struct {
    state int
}

func (s *StreamService) render_router(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*83) % 997
    }
    return value
}

func main() {
    obj := &StreamService{state: 83}
    fmt.Println(obj.render_router(83))
}
