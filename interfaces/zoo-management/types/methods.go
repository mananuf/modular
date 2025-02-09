package types

import (
	"fmt"
	"time"
)

func (l *Lion) MakeSound() string {
	return "roar"
}
func (l *Lion) Eat() string {
	fmt.Println("------->",l.LastFeedTime)

	if l.LastFeedTime != "" {
		lastFeed, err := time.Parse(defaultDate, l.LastFeedTime)

		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}

		if time.Since(lastFeed).Hours() < 4 {
			return fmt.Sprintf("%s has already been feed at %s", l.Name, l.LastFeedTime)
		}
	}

	l.FoodType = MEAT
	l.LastFeedTime = time.Now().Format(defaultDate)

	return fmt.Sprintf("%s just got fed @ %s", l.Name, l.LastFeedTime)
}

func (l *Lion) Sleep() string {
	l.IsSleeping = true
	return "Lion was put to sleep"
}

func (ld *LionsDen) AddLionToDen(lion Lion) {
	ld.AllLions = append(ld.AllLions, lion)
}

func (e *Elephant) MakeSound() string {
	return "rumble"
}

func (e *Elephant) Eat() string {
	e.FoodType = OMNIVORE
	return "Elephant was just fed"
}

func (e *Elephant) Sleep() string {
	e.IsSleeping = true
	return "Elephant was put to sleep"
}

func (m *Monkey) MakeSound() string {
	return "howl"
}

func (m *Monkey) Eat() string {
	return "Monkey was just fed"
}

func (m *Monkey) Sleep() string {
	m.IsSleeping = true
	return "Monkey was put to sleep"
}

func (z *ZooKeeper) AddAnimalsToZooKeeper(animals ...Animal) {
	var animalsToCaterFor []Animal
	animalsToCaterFor = append(animalsToCaterFor, animals...)
	z.Animal = animalsToCaterFor

	fmt.Printf("Animals to cater for: %v", animalsToCaterFor)
}

func (z *ZooKeeper) FeedAnimal(animal Animal) {
	fmt.Println(animal.Eat())
}

// --------------------------------------------------------
// --------------- constructor methods --------------------
// --------------------------------------------------------

func NewAnimalInfo(name string, weight float64, foodType FoodType) *AnimalInfo {
	return &AnimalInfo{
		Name:         name,
		Weight:       weight,
		IsSleeping:   false,
		FoodType:     foodType,
		LastFeedTime: "",
	}
}

func NewLionsDen() *LionsDen {
	return &LionsDen{
		AllLions: make(Lions, 0),
	}
}

func NewLion(info AnimalInfo) *Lion {
	return &Lion{info}
}

func NewElephant(info AnimalInfo) *Elephant {
	return &Elephant{info}
}

func NewMonkey(info AnimalInfo) *Monkey {
	return &Monkey{info}
}

func NewZooKeeper(name string, shift string) *ZooKeeper {
	return &ZooKeeper{
		Name:   name,
		Shift:  shift,
		Animal: make([]Animal, 0),
	}
}
