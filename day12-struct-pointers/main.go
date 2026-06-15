package main

import "fmt"

type Server struct {
	Name   string
	CPU    int
	Memory int
}

func (s *Server) increaseCPU() {
	s.CPU = s.CPU + 10
}

func (s *Server) increaseMemory() {
	s.Memory = s.Memory + 4
}

func (s *Server) restart() {
	s.CPU = 0
	s.Memory = 0
}

func main() {
	server := Server{
		Name:   "Production-1",
		CPU:    50,
		Memory: 16,
	}

	server.increaseCPU()
	server.increaseMemory()
	server.restart()

	fmt.Println("CPU: ", server.CPU)
	fmt.Println("Memory: ", server.Memory)
}
