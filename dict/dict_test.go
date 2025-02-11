package dict

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	m1 := map[string]string{}
	m1["key1"] = "value1"
	assert.True(t, Contains(m1, "key1"))
	assert.False(t, Contains(m1, "key2"))

	var m2 map[string]string = nil
	assert.False(t, Contains(m2, "key1"))

	var m3 map[string]string = make(map[string]string)
	assert.False(t, Contains(m3, "key1"))
}

func TestKeys(t *testing.T) {
	m1 := map[string]string{}
	m1["key3"] = "value3"
	m1["key2"] = "value2"
	m1["key1"] = "value1"
	keys := Keys(m1)
	assert.Equal(t, 3, len(keys))
	sort.Strings(keys)
	assert.Equal(t, []string{"key1", "key2", "key3"}, keys)

	var m2 map[string]string = nil
	keys = Keys(m2)
	assert.Equal(t, 0, len(keys))

	var m3 map[string]string = make(map[string]string)
	keys = Keys(m3)
	assert.Equal(t, 0, len(keys))
}

func TestValues(t *testing.T) {
	m1 := map[string]string{}
	m1["key3"] = "value3"
	m1["key2"] = "value2"
	m1["key1"] = "value1"
	values := Values(m1)
	assert.Equal(t, 3, len(values))
	sort.Strings(values)
	assert.Equal(t, []string{"value1", "value2", "value3"}, values)

	var m2 map[string]string = nil
	values = Values(m2)
	assert.Equal(t, 0, len(values))

	var m3 map[string]string = make(map[string]string)
	values = Values(m3)
	assert.Equal(t, 0, len(values))
}
