package wakeup

import "fmt"

type Bag struct {
	Matches bool
	Torch   bool
	Knife   bool
}
type Character struct {
	Name string
	Bag  Bag
}

func Wakeup(c Character) {
	items := ""
	if c.Bag.Matches {
		items += "Сірники "
	}
	if c.Bag.Torch {
		items += "Ліхтарик "
	}
	if c.Bag.Knife {
		items += "Ніж "
	} else {
		items += "Нічого"
	}

	fmt.Printf("%s прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить: %s\n", c.Name, items)
}
