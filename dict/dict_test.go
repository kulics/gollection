package dict

import (
	"fmt"
	"testing"
)

func TestHashDict(t *testing.T) {
	var dict1 = Of[string, int]()
	if dict1.Count() != 0 {
		t.Fatal("dict count not eq 0")
	}
	dict1.Add("111", 1)
	if dict1.Count() != 1 {
		t.Fatal("dict count not eq 1")
	}
	dict1.Add("111", 2)
	if dict1.Count() != 1 {
		t.Fatal("dict count not eq 1")
	}
	if v, ok := dict1.At("111").Val(); !ok || v != 2 {
		t.Fatal("dict value not eq 2")
	}
	dict1.At("111").Set(3)
	if dict1.At("111").Get() != 3 {
		t.Fatal("dict value not eq 3")
	}
	var strkey = fmt.Sprintf("%d", 111)
	dict1.Add(strkey, 3)
	if dict1.Count() != 1 {
		t.Fatal("dict count not eq 1")
	}
	if v, ok := dict1.At(strkey).Val(); !ok || v != 3 {
		t.Fatal("dict value not eq 3")
	}
	var dict2 = Of[int, int]()
	if dict2.Count() != 0 {
		t.Fatal("dict count not eq 0")
	}
	dict2.Add(111, 1)
	if dict2.Count() != 1 {
		t.Fatal("dict count not eq 1")
	}
	dict2.Add(111, 2)
	if dict2.Count() != 1 {
		t.Fatal("dict count not eq 1")
	}
	if v, ok := dict2.At(111).Val(); !ok || v != 2 {
		t.Fatal("dict value not eq 2")
	}
}
