package main

import (
	"fmt"

	"quizz/character"
	"quizz/story"
)

func main() {
	fmt.Printf("Гра розпочинається: \n")

	Stephen := character.Character{
		Name: "Стівен",
		Bag: character.Bag{
			Matches: true,
			Torch:   true,
			Knife:   true,
		},
	}

	story.Wakeup(Stephen)
	story.Meetanimal(Stephen)
	story.Findbox(Stephen)
}
