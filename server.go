package main

import "fmt"

type JunioQueue struct {
	Project   string
	Name      string
	NextIndex int
}

func NewJunioQueue() JunioQueue {
	return JunioQueue{}
}

func main() {
	fmt.Println("starting junio server")
}
