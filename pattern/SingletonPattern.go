package pattern

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type config struct {
	// config variables
}

var counter int = 1
var singleConfigInstance *config

// GetConfigInstance is a function to get singleton instance
func GetConfigInstance() *config {
	if singleConfigInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleConfigInstance == nil {
			fmt.Println("Creating Single Instance Now!, and counter is ", counter)
			singleConfigInstance = &config{}
			counter++
		} else {
			fmt.Println("First Single Instance already Created, returning it!, and counter is ", counter)
		}
	} else {
		fmt.Println("Second Single Instance already Created, returning it!, and counter is ", counter)
	}
	return singleConfigInstance
}

func GetSingletonName(s string) string {
	return "Code for : -- " + s + " Pattern --"
}
