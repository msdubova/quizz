package character

type Bag struct {
	Matches bool
	Torch   bool
	Knife   bool
}
type Character struct {
	Name string
	Bag  Bag
}
