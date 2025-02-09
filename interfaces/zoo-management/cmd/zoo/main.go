package main

import (
	"fmt"
	"time"

	"github.com/mananuf/zoo-management/types"
)

func main() {
	lionInfo := types.NewAnimalInfo(
		"simba",
		40.34,
		types.MEAT,
	)

	simba := types.NewLion(*lionInfo)
	ld := types.NewLionsDen()

	ld.AddLionToDen(*simba)
	
	zk := types.NewZooKeeper(
		"shokkz",
		time.Now().Format(time.Kitchen),
	)

	zk.AddAnimalsToZooKeeper(simba)

	fmt.Println(zk)

	zk.FeedAnimal(simba)

	fmt.Println(simba)

	zk.FeedAnimal(simba)
}