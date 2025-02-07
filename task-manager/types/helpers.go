package types

import "fmt"

func PrintActualValues(allTasks []*Task) {
	for _, task := range allTasks {
		fmt.Println(task)
	}
}