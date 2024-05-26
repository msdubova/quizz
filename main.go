package main

import (
	"fmt"

	"quizz/meetanimal"
	"quizz/wakeup"
)

func main() {
	fmt.Printf("Game started: \n")

	Stephen := wakeup.Character{
		Name: "Стівен",
		Bag: wakeup.Bag{
			Matches: true,
			Torch:   true,
			Knife:   true,
		},
	}

	wakeup.Wakeup(Stephen)
	meetanimal.Meetanimal(Stephen)
}
