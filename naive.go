package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func concat(word string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += word
	}
	return s
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("./naive <word> <n> <repeat>")
	}

	word := os.Args[1]
	n, _ := strconv.ParseInt(os.Args[2], 10, 64)
	repeat, _ := strconv.ParseInt(os.Args[3], 10, 64)

	var avg time.Duration
	for i := 0; i < int(repeat); i++ {
		start := time.Now()
		concat(word, int(n))
		elapsed := time.Now().Sub(start)
		avg += elapsed
	}

	result := avg.Seconds() / float64(repeat)
	fmt.Println(result)
}
