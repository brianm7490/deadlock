package main

import "fmt"

type BatchResolver struct {
    state int
}

func (s *BatchResolver) compute_loader(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*83) % 997
    }
    return result
}

func main() {
    obj := &BatchResolver{state: 83}
    fmt.Println(obj.compute_loader(83))
}
