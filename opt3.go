package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func concat(word []byte, n int) []byte {
	var sb []byte

	for i := 0; i < n; i++ {
		sb = append(sb, word...)
	}
	return sb
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("./opt3 <word> <n> <repeat>")
	}

	word := []byte(os.Args[1])
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
