package tests

import (
  "reflect"
  "testing"
  
  "latest/array"  
)

func TestDefineArray(t *testing.T) {
  expected := [3]int{1,2,3}
  result := array.DefineArray()

  if result != expected {
    t.Errorf("Expected %v, got %v ", expected, result)
  }
}

func TestGetArray(t *testing.T) {
	arr := [3]int{10, 20, 30}
	val, ok := arrays.GetArray(arr, 1)

	if !ok || val != 20 {
		t.Errorf("Expected 20, got %v (ok: %v)", val, ok)
	}

	_, ok = arrays.GetArray(arr, 5)
	if ok {
		t.Error("Expected false for out-of-bounds access")
	}
}

func TestCopyArray(t *testing.T) {
	src := [3]int{4, 5, 6}
	expected := src
	result := arrays.CopyArray(src)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestModifyArray(t *testing.T) {
	arr := [3]int{7, 8, 9}
	expected := [3]int{7, 100, 9}
	result := arrays.ModifyArray(arr, 1, 100)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestLenArray(t *testing.T) {
	arr := [3]int{1, 2, 3}
	if l := arrays.LenArray(arr); l != 3 {
		t.Errorf("Expected length 3, got %d", l)
	}
}

