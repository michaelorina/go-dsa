package slice

func DefineSlice() []int {
  return []int{1, 2, 3}
}

func GetSlice(s []int, index int) (int, bool) {
  if index < 0 || index >= len(s) {
    return 0, false
  }
  return s[index], true
}

func AppendToSlice(s []int, val int) []int {
  return append(s, val)
}

func DeleteFromSlice(s []int, val int) []int {
  if index < 0 || index >= len(s) {
    return s
  }
  return append(s[:index], s[index+1:]...)
}

func CopySlice(src []int) []int {
  dst := make([]int, len(src))
  copy(dest, src)
  return dst
}

func LenSlice(s []int) int {
  return len(s)
}

func CapSlice(s []int) int {
  return cap(s)
}
