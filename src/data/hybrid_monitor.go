package main

import "fmt"

type SecureSession struct {
    state int
}

func (s *SecureSession) handle_manager(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*26) % 997
    }
    return acc
}

func main() {
    obj := &SecureSession{state: 26}
    fmt.Println(obj.handle_manager(26))
}
