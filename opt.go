package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type StringBuilder struct {
	buff []string
}

func (sb *StringBuilder) append(word string) *StringBuilder {
	sb.buff = append(sb.buff, word)
	return sb
}

func (sb *StringBuilder) build() string {
	return strings.Join(sb.buff, "")
}

func concat(word string, n int) string {
	sb := StringBuilder{
		buff: make([]string, 0),
	}

	for i := 0; i < n; i++ {
		sb.append(word)
	}
	return sb.build()
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("./opt <word> <n> <repeat>")
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
