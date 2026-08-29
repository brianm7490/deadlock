package main

import "fmt"

type AtomicContext struct {
    state int
}

func (s *AtomicContext) render_loader(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*47) % 997
    }
    return value
}

func main() {
    obj := &AtomicContext{state: 47}
    fmt.Println(obj.render_loader(47))
}
