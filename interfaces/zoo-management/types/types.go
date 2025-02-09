package types

type FoodType string
type Lions []Lion
type Elephants []Elephant
type Monkies []Monkey


type Animal interface {
	Eat() string
	Sleep() string
	MakeSound() string
}

type AnimalInfo struct {
	Name         string
	Weight       float64
	FoodType     FoodType
	IsSleeping   bool
	LastFeedTime string
}

type Lion struct {
	AnimalInfo
}

type Elephant struct {
	AnimalInfo
}

type Monkey struct {
	AnimalInfo
}

type LionsDen struct {
	AllLions Lions
}

type ZooKeeper struct {
	Name string
	Shift string
	Animal []Animal
}
