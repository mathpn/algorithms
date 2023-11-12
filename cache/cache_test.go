package cache

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	type test struct {
		op    string
		key   string
		value int
		err   bool
	}

	tests := []test{
		{op: "get", key: "foo", value: 3, err: true},
		{op: "update", key: "foo", value: 3, err: false},
		{op: "get", key: "foo", value: 3, err: false},
		{op: "update", key: "bar", value: 42, err: false},
		{op: "get", key: "bar", value: 42, err: false},
		{op: "update", key: "ping", value: 11, err: false},
		{op: "get", key: "ping", value: 11, err: false},
		{op: "update", key: "pong", value: 987, err: false},
		{op: "get", key: "pong", value: 987, err: false},
		{op: "get", key: "foo", value: 3, err: true},
		{op: "get", key: "bar", value: 42, err: false},
		{op: "update", key: "foo", value: 3, err: false},
		{op: "get", key: "bar", value: 42, err: false},
		{op: "get", key: "foo", value: 3, err: false},
		{op: "get", key: "ping", value: 11, err: true},
	}

	lru := NewLRU[string, int](3)

	for i, test := range tests {
		if test.op == "update" {
			lru.Update(test.key, test.value)
		} else {
			v, err := lru.Get(test.key)
			if test.err && err == nil {
				fmt.Println(v)
				t.Errorf("operation %d: expected error when getting key %s", i, test.key)
			}
			if !test.err && err != nil {
				t.Fatalf("operation %d: %s", i, err)
			}
			if !test.err && v != test.value {
				t.Errorf("operation %d: expected value %d, got %d", i, test.value, v)
			}
		}
	}
}
