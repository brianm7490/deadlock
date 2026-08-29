package main

import "fmt"

type AtomicCache struct {
    state int
}

func (s *AtomicCache) fetch_provider(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*13) % 997
    }
    return acc
}

func main() {
    obj := &AtomicCache{state: 13}
    fmt.Println(obj.fetch_provider(13))
}
