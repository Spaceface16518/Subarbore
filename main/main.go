package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"subarbore"
)

var input string
var regular bool
var unchecked bool

func init() {
	log.SetOutput(os.Stderr)
	log.SetPrefix("subarbore> ")

	flag.StringVar(&input, "input", "", "Use command line input instead of reading from standard input")
	flag.BoolVar(&regular, "regular", false, "Use regular notation instead of subtractive notation")
	flag.BoolVar(&unchecked, "unchecked", false, "Do not validate string. This will speed up translation but will yield wrong answers without warning if input is invalid")

	flag.Parse()
}

func main() {
	if regular {
		subarbore.SetGlobalDefaultRegular()
	}
	var text string
	if input == "" {
		reader := bufio.NewReader(os.Stdin)
		var err error
		text, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalf("An error was encountered while reading from standard input: %v\n", err)
		}
	} else {
		text = input
	}
	output := process(text)
	writeOut(output)
}

func process(s string) int {
	if unchecked {
		return subarbore.Parse(s)
	}
	decimal, invalid := subarbore.ParseChecked(s)
	if invalid != nil {
		log.Fatalf("The numeral provided was invalid. The offensive numerals were found to numeral group(s) %s", formatIntList(invalid))
	}
	return decimal
}

func formatIntList(l []int) string {
	length := len(l)
	if length == 0 {
		return ""
	} else if length == 1 {
		return fmt.Sprintf("%d", l[0])
	} else if length == 2 {
		return fmt.Sprintf("%d and %d", l[0], l[1])
	} else {
		return fmt.Sprintf("%d, ", l[0]) + formatIntList(l[1:])
	}
}

func writeOut(i int) {
	fmt.Printf("%d\n", i)
}
