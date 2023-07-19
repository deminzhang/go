package util

func SliceFind[T comparable](slice []T, dst T) int {
	for i, v := range slice {
		if v == dst {
			return i
		}
	}
	return -1 // not found
}

func SliceRemove[T any](slice []T, index int) []T {
	lastIdx := len(slice) - 1
	if lastIdx < 0 || index < 0 || index > lastIdx {
		return slice
	}

	if index != lastIdx {
		slice[index] = slice[lastIdx]
	}
	return slice[:lastIdx]
}

func SliceRemoveWithOrder[T any](slice []T, index int) []T {
	if index >= 0 && index < len(slice) {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}
