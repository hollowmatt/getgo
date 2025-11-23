package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ch 7")
	var store Store
	err := store.Put("key1", "value1")
	if err != nil {
		fmt.Println("Error putting value:", err)
		return
	}
}

type Store interface {
	Put(key string, val string) error
	Get(key string) (string, error)
	Delete(key string) error
}
