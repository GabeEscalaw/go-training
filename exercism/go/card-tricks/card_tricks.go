package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if len(slice) > 0 && index < len(slice) && index >= 0{
		return slice[index], true
	} else {
		return 0, false
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice) > 0 && index < len(slice) && index >= 0{
		slice[index] = value
		return slice
	} else {
		return append(slice, value)
	}
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var s []int

	if length > 0 {
		for i := 0; i < length; i++ {
			s = append(s, value)
		}
	} 
	
	return s
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice) > 0 && index == 0 { 
		slice = slice[1:] 
	} else if len(slice) > 0 && index == len(slice)-1 { 
		slice = slice[:index] 
	} else if len(slice) > 0 && index < len(slice)-1 && index >= 0{
		slicedMid := append(slice[:index], slice[index+1:]...)
		slice = slicedMid
	}

	return slice
}
