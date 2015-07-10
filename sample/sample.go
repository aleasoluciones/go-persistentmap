package main

import (
	"fmt"

	"encoding/json"
	"github.com/aleasoluciones/go-persistentmap"
)

type Dummy struct {
	Id   string
	Name string
}

func dummySerialize(obj interface{}) []byte {
	serilized, _ := json.Marshal(obj)
	return serilized
}

func dummyDeserialize(serialized []byte) interface{} {
	obj := Dummy{}
	if err := json.Unmarshal(serialized, &obj); err != nil {
		fmt.Println("Deserialization error", err, serialized)
	}
	return obj
}

func main() {

	m := persistentmap.NewPersistentMap("test.db")
	fmt.Println(m)

	m.Set("answer1", []byte("42"))
	m.Set("answer2", []byte("43"))
	m.Set("answer3", []byte("44"))

	for tuple := range m.IterationChannel() {
		fmt.Println("Tuple", tuple.Key, string(tuple.Value))
	}

	fmt.Println("E1", string(m.Get("answer1")))
	m.Delete("answer1")

	fmt.Println("E2", string(m.Get("answer1")))

	m2 := persistentmap.NewPersistentMapWithSerialization("test2.db", dummySerialize, dummyDeserialize)
	m2.SerializeAndSet("id1", Dummy{"id1", "john"})
	m2.SerializeAndSet("id2", Dummy{"id2", "kent"})

	result := m2.GetAndDeserialize("id1")
	fmt.Println("Deserialized %s", result, fmt.Sprintf("%T", result))
	for tuple := range m2.IterationDeserializedChannel() {
		fmt.Println("Tuple", tuple.Key, tuple.Value, fmt.Sprintf("%T", tuple.Value))
	}
}
