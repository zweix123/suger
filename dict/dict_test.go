package dict

import (
	"sort"
	"testing"
)

func TestContain(t *testing.T) {
	m1 := map[string]string{}
	m1["key1"] = "value1"
	if !Contain(m1, "key1") {
		t.Errorf("Contain should return true")
	}
	if Contain(m1, "key2") {
		t.Errorf("Contain should return false")
	}

	var m2 map[string]string = nil
	if Contain(m2, "key1") {
		t.Errorf("Contain should return false")
	}

	var m3 map[string]string = make(map[string]string)
	if Contain(m3, "key1") {
		t.Errorf("Contain should return false")
	}
}

func TestKeys(t *testing.T) {
	m1 := map[string]string{}
	m1["key3"] = "value3"
	m1["key2"] = "value2"
	m1["key1"] = "value1"
	keys := Keys(m1)
	if len(keys) != 3 {
		t.Errorf("Keys should return 3 keys")
	}
	sort.Strings(keys)
	if keys[0] != "key1" || keys[1] != "key2" || keys[2] != "key3" {
		t.Errorf("Keys should return key1, key2, key3")
	}

	var m2 map[string]string = nil
	keys = Keys(m2)
	if len(keys) != 0 {
		t.Errorf("Keys should return 0 keys")
	}

	var m3 map[string]string = make(map[string]string)
	keys = Keys(m3)
	if len(keys) != 0 {
		t.Errorf("Keys should return 0 keys")
	}
}

func TestValues(t *testing.T) {
	m1 := map[string]string{}
	m1["key3"] = "value3"
	m1["key2"] = "value2"
	m1["key1"] = "value1"
	values := Values(m1)
	if len(values) != 3 {
		t.Errorf("Values should return 3 values")
	}
	sort.Strings(values)
	if values[0] != "value1" || values[1] != "value2" || values[2] != "value3" {
		t.Errorf("Values should return value1, value2, value3")
	}

	var m2 map[string]string = nil
	values = Values(m2)
	if len(values) != 0 {
		t.Errorf("Values should return 0 values")
	}

	var m3 map[string]string = make(map[string]string)
	values = Values(m3)
	if len(values) != 0 {
		t.Errorf("Values should return 0 values")
	}
}
