package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name   string `json:"name"`
	CPU    int    `json:"cpu"`
	Status string `json:"status"`
	Memory int    `json:"memory"`
}

func main() {
	servers := []Server{

		{
			Name:   "Production-1",
			CPU:    70,
			Status: "Running",
			Memory: 16,
		},

		{
			Name:   "Production-2",
			CPU:    90,
			Status: "High CPU",
			Memory: 32,
		},

		{
			Name:   "Production-3",
			CPU:    50,
			Status: "Medium Load",
			Memory: 8,
		},
	}

	jsonData, err := json.MarshalIndent(servers, "", "  ")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
