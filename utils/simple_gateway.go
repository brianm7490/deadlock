package main

import "fmt"

type SharedLoader struct {
    state int
}

func (s *SharedLoader) fetch_monitor(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*39) % 997
    }
    return value
}

func main() {
    obj := &SharedLoader{state: 39}
    fmt.Println(obj.fetch_monitor(39))
}
