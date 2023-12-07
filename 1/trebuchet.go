package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("trebuchet go")
	numsFromLines := []int{}
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		digitsFromLine := []int{}
		line = scanner.Text()
		//digitsFromLine is all digits in a line of input text
		for _, c := range line {
			if isNumeric(string(c)) {
				//c-'0' converts rune ASCII to int value
				digitsFromLine = append(digitsFromLine, int(c-'0'))
			}
		}
		//numsFromLines will be first and last digit from each line
		if len(digitsFromLine) == 1 {
			numToAdd := digitsFromLine[0]*10 + digitsFromLine[0]
			numsFromLines = append(numsFromLines, numToAdd)
		} else if len(digitsFromLine) == 2 {
			numToAdd := digitsFromLine[0]*10 + digitsFromLine[1]
			numsFromLines = append(numsFromLines, numToAdd)
		} else if len(digitsFromLine) > 2 {
			z := len(digitsFromLine) - 1
			numToAdd := digitsFromLine[0]*10 + digitsFromLine[z]
			numsFromLines = append(numsFromLines, numToAdd)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sum := 0
	fmt.Printf("%d numbers to add\n", len(numsFromLines))
	for _, num := range numsFromLines {
		sum = num + sum
	}
	fmt.Printf("final sum: %d\n", sum)

	duration := time.Since(start)
	fmt.Println("finished in", duration)
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 32)
	return err == nil
}
