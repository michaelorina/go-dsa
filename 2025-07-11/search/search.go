package search

func LinearSearch(array []int, target int) (int, bool) {
  for index, value := range array {
    if value == target {
      return index, true
    }
  }
  return -1, false
}

func BinarySearch(array []int, target int) (int, bool) {
  low, high := 0, len(array)

  for low < high {
    mid := low + (high - low) / 2

    if array[mid] == target {
      return mid, true
    } else if array[mid] < target {
      low = mid + 1
    } else {
      high = mid
    }
  }
  return -1, false
}
