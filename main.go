package main

import (
	"fmt"

	"quizz/findbox"
	"quizz/meetanimal"
	"quizz/wakeup"
)

func main() {
	fmt.Printf("Гра розпочинається: \n")

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
	findbox.Findbox(Stephen)
}
