package tests

import (
  "reflect"
  "testing"

  "github.com/michaelorina/go-dsa/latest/slice"
)

func TestDefineSlice(t *testing.T) {
  expected := []int{1,2,3}
  result := slice.DefineSlice()

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestGetSlice(t *testing.T) {
  s := []int{10, 20, 30}
  val, ok := slice.GetSlice(s, 2)

  if !ok || val != 30 {
    t.Errorf("Expected 30, got %v (ok: %v)", val, ok)
  }

  _, ok = slice.GetSlice(s, 5)
    if ok {
      t.Errorf("Expected false for out-of-bounds access")
    }
}

func TestAppendToSlice(t *testing.T) {
  s := []int{1, 2}
  expected := []int{1, 2, 3}
  result := slice.AppendToSlice(s, 3)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestDeleteFromSlice(t *testing.T) {
  s := []int{1, 2, 3, 4}
  expected := []int{1, 3, 4}
  result := slice.DeleteFromSlice(s, 1)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestCopySlice(t *testing.T) {
  s := []int{4, 5, 6}
  expected := []int{4, 5, 6}
  result := slice.CopySlice(s)
  
  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }

  if &result[0] == &s[0] {
    t.Errorf("Expected a deep copy, but elements share memory")
  }
}

func TestLenSlice(t *testing.T){
  s := []int{1, 2, 3, 4}
  if l := slice.LenSlice(s); l != 4 {
    t.Errorf("Expected length 4, got %d", l)
  }  
}

func TestCapSlice(t *testing.T) {
  s := make([]int, 2, 5)
  if c := slice.CapSlice(s); c != 5 {
    t.Errorf("Expected capacity 5, got %d", c)
  }
}
