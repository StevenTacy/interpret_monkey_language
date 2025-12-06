package object

import "testing"

func TestStringHashedKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("hello1 and hello2 has same content but got different hash key")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("diff1 and diff2 has same content but got different hash key")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("hello1 and diff1 has different content but got same hash key")
	}
}
