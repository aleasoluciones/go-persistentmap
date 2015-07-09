package main

import (
	"fmt"

	"github.com/aleasoluciones/go-persistentmap"
)

func main() {

	persistentmap := persistentmap.NewPersistentMap("test.db")
	fmt.Println(persistentmap)

	persistentmap.Set("answer1", []byte("42"))
	persistentmap.Set("answer2", []byte("43"))
	persistentmap.Set("answer3", []byte("44"))

	for tuple := range persistentmap.IterationChannel() {
		fmt.Println("Tuple", tuple.Key, string(tuple.Value))
	}

	fmt.Println("E1", string(persistentmap.Get("answer1")))
	persistentmap.Delete("answer1")

	fmt.Println("E2", string(persistentmap.Get("answer1")))

}
