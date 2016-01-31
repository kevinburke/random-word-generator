package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
	"unicode"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func startsWithUppercase(word string) bool {
	for _, runeValue := range word {
		return unicode.IsUpper(runeValue)
	}
	return false
}

func main() {
	var count = flag.Int("count", 1, "Number of words to print")
	flag.Parse()

	bits, err := ioutil.ReadFile("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	words := bytes.Split(bits, []byte("\n"))
	for i := 0; i < *count; i++ {
		for {
			randomVal := r.Intn(len(words))
			word := string(words[randomVal])
			if startsWithUppercase(word) {
				continue
			} else {
				fmt.Println(word)
				break
			}
		}
	}
}
