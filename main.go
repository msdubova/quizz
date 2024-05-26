package main

import (
	"fmt"

	"quizz/wakeup"
)

func main() {
	fmt.Printf("Game started: ")

	Stephen := wakeup.Character{
		Name: "Стівен",
		Bag: wakeup.Bag{
			Matches: true,
			Torch:   true,
			Knife:   true,
		},
	}

	wakeup.Wakeup(Stephen)
}
