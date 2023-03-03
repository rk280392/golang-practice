package main

import (
	"fmt"
	"net"
	"sort"
	"strings"
	"time"
)

func printProgressBar(iteration, total int, prefix, suffix string, length int, fill string) {
	percent := int64(float64(iteration) / float64(total) * 100)
	filledLength := int(length * iteration / total)
	end := ">"
	if iteration == total {
		end = "="
	}
	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
	fmt.Printf("\r%s [%s] %d%% %s", prefix, bar, percent, suffix)
	if iteration == total {
		fmt.Println()
	}
}

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 300)
	results := make(chan int)
	var openports []int
	totalPorts := 1024

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= totalPorts; i++ {
			ports <- i
		}
	}()

	for i := 0; i < totalPorts; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
		time.Sleep(100 * time.Millisecond) // mimics work
		printProgressBar(i+1, totalPorts, "Progress", "Complete", 100, "=")
	}
	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
