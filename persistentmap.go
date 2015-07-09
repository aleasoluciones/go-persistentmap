package persistentmap

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const (
	MapBucket = "map"
)

type SerializeFunc func(interface{}) []byte
type DeserializeFunc func([]byte) interface{}

type PersistentMap struct {
	db           *bolt.DB
	name         string
	serializer   SerializeFunc
	deserializer DeserializeFunc
}

func NewPersistentMap(filename string) *PersistentMap {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(MapBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return &PersistentMap{db: db, name: MapBucket}
}

func NewPersistentMapWithSerialization(filename string, serializer SerializeFunc, deserializer DeserializeFunc) *PersistentMap {
	m := NewPersistentMap(filename)
	m.serializer = serializer
	m.deserializer = deserializer
	return m
}

func (m *PersistentMap) SerializeAndSet(key string, obj interface{}) {
	m.Set(key, m.serializer(obj))
}

func (m *PersistentMap) GetAndDeserialize(key string) interface{} {
	return m.deserializer(m.Get(key))
}

func (m *PersistentMap) Set(key string, data []byte) {
	m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(m.name))
		err := b.Put([]byte(key), data)
		return err
	})
}

func (m *PersistentMap) Get(key string) []byte {
	returnValue := []byte{}
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(m.name))
		returnValue = b.Get([]byte(key))
		return nil
	})
	return returnValue
}

func (m *PersistentMap) Close() {
	m.db.Close()
}

func (m *PersistentMap) Delete(key string) {
	m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(m.name))
		err := b.Delete([]byte(key))
		return err
	})
}

type Tuple struct {
	Key   string
	Value []byte
}

func (m *PersistentMap) IterationChannel() chan Tuple {
	returnChan := make(chan Tuple)

	go func() {
		m.db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(m.name))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				returnChan <- Tuple{string(k), v}
			}
			close(returnChan)
			return nil
		})
	}()

	return returnChan
}
