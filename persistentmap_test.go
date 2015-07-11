// +build integration

package persistentmap_test

import (
	"fmt"
	"testing"

	. "github.com/aleasoluciones/go-persistentmap"
	"github.com/stretchr/testify/assert"
)

func TestRawBytesMap(t *testing.T) {
	t.Parallel()

	m := NewPersistentMap("rawmap1.db")
	fmt.Println(m)

	m.Set("k1", []byte("value1"))
	m.Set("k2", []byte("value2"))

	assert.Equal(t, m.Get("k1"), []byte("value1"))
	assert.Equal(t, m.Get("k2"), []byte("value2"))
}

func TestRawBytesMapIteration(t *testing.T) {
	t.Parallel()

	m := NewPersistentMap("rawmap2.db")
	fmt.Println(m)

	m.Set("k1", []byte("value1"))
	m.Set("k2", []byte("value2"))

	c := m.IterationChannel()
	fmt.Println("EFA1")

	assert.Equal(t, <-c, Tuple{"k1", []byte("value1")})
	fmt.Println("EFA2")
	assert.Equal(t, <-c, Tuple{"k2", []byte("value2")})
}
