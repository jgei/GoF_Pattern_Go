package singleton_by_mutex_lib

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {
}

var singleInstance *singleton

func GetInstance() *singleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &singleton{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
