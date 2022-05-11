//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups
package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Count struct {
	count int
	sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	counter := Count{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wg.Add(1)
		fmt.Println(scanner.Text())
		go func(str string) {
			counter.Lock()
			defer counter.Unlock()
			defer wg.Done()
			counter.count += len([]rune(str))
		}(scanner.Text())
	}

	wg.Wait()
	counter.Lock()
	sum := counter.count
	counter.Unlock()

	fmt.Println("Total number of letters:", sum)
}
