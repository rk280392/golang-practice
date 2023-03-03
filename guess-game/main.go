// guess challenges players to guess a random number

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("We have chosen a number from 1 to 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have 10 guesses")
		fmt.Print("Make a guess : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your guess was LOW")
		} else if guess > target {
			fmt.Println("Your guess was High")
		} else {
			success = true
			fmt.Println("Congratulations!! you guessed the right number!!")
			break
		}
	}

	if !success {
		fmt.Printf("The right number was : %d\n", target)
	}

}
