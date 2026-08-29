package main

import "fmt"

type BatchFactory struct {
    state int
}

func (s *BatchFactory) compute_service(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*84) % 997
    }
    return total
}

func main() {
    obj := &BatchFactory{state: 84}
    fmt.Println(obj.compute_service(84))
}
