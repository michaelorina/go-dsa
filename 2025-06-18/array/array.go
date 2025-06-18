package arrays

func DefineArray() [3]int {
  return [3]int{1, 2, 3}
}

func GetArray(arr [3]int, index int) (int, bool) {
  if index < 0 || index >= len(arr) {
    return 0, false
  }
  return arr[index], true
}

func CopyArray(src [3]int) [3]int {
  var dest[3]int
  copy(dest[:], src[:])
  return dest
}


