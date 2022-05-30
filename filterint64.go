package filterint64

func Where(values []int64, callback func(value int64) bool) []int64 {
	var result []int64

	for _, value := range values {
		if callback(value) {
			result = append(result, value)
		}
	}

	return result
}
