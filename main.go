package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	names := []string{"web-frontend", "api-server", "worker-1", "redis-cache", "postgres-db"}
	images := []string{"nginx:alpine", "node:18-slim", "python:3.11", "redis:7", "postgres:15"}
	fmt.Println("Container Orchestrator Status")
	fmt.Println("=============================")
	running, stopped := 0, 0
	for i := range names {
		icon := "UP"
		if rand.Float64() < 0.1 {
			icon = "!!"
			stopped++
		} else {
			running++
		}
		ports := fmt.Sprintf("%d:80", 8080+i)
		fmt.Printf("  [%s] %-18s %-20s %s\n", icon, names[i], images[i], ports)
	}
	fmt.Printf("\n%d running, %d issues, %d total\n", running, stopped, len(names))
}
