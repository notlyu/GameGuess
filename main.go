package main

import (
	"fmt"
	"math/rand/v2"
)

type GameGuess struct {
	lvl      int
	attempts int
	guess    int
}

func (choice *GameGuess) ChoiceLvl() {

	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	for {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice.lvl)

		switch choice.lvl {
		case 1:
			choice.attempts = 10
			return
		case 2:
			choice.attempts = 5
			return
		case 3:
			choice.attempts = 3
			return
		default:
			fmt.Println("Invalid choice. Enter 1, 2, or 3.")

		}
	}

}

func (g *GameGuess) Play() {

	target := rand.IntN(100) + 1
	won := false
	for attempts := 1; attempts <= g.attempts; attempts++ {

		fmt.Print("Enter your guess: ")
		fmt.Scan(&g.guess)

		if g.guess == target {
			won = true
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			break
		}
		if g.guess < target {
			fmt.Printf("Incorrect! The number is greater than %d.\n", g.guess)
		} else {
			fmt.Printf("Incorrect! The number is less than %d.\n", g.guess)
		}

	}
	if !won {
		fmt.Printf("You ran out of attempts! The number was %d.\n", target)
	}

}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	c := GameGuess{}

	c.ChoiceLvl()

	var name string
	switch c.lvl {
	case 1:
		name = "Easy"
	case 2:
		name = "Medium"
	case 3:
		name = "Hard"
	}
	fmt.Printf("Great! You have selected the %s difficulty level.\n", name)
	fmt.Println("Let's start the game!")
	fmt.Println("")

	c.Play()
}
