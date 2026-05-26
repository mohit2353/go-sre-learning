package main

import "fmt"

type Server struct {
	Name   string
	CPU    int
	Memory int
}

func (s Server) checkHealth() {

	fmt.Println("Server:", s.Name)

	if s.CPU > 80 {
		fmt.Println("High CPU usage")
	} else {
		fmt.Println("CPU Healthy")
	}

	if s.Memory > 80 {
		fmt.Println("High Memory usage")
	} else {
		fmt.Println("Memory Healthy")
	}
}

func (s Server) display() {
	fmt.Println("Server:", s.Name)
	fmt.Println("CPU:", s.CPU)
	fmt.Println("Memory:", s.Memory)
}

func main() {
	server1 := Server{
		Name:   "Production-1",
		CPU:    75,
		Memory: 90,
	}
	server1.checkHealth()
	server1.display()

}
