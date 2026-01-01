package refutils

func ToPointers[T any](boxes []T) []*T {
	result := make([]*T, len(boxes))
	for i := range boxes {
		result[i] = &boxes[i]
	}
	return result
}
