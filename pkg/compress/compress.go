package compress

// Compression compresses example string input
func Compress(exStr string) []int {
	code := 256
	mmap := make(map[string]int)
	for i := 0; i < 256; i++ {
		mmap[string(i)] = i
	}

	currentChar := ""
	result := make([]int, 0)
	for _, c := range []byte(exStr) {
		phrase := currentChar + string(c)
		if _, isTrue := mmap[phrase]; isTrue {
			currentChar = phrase
		} else {
			result = append(result, mmap[currentChar])
			mmap[phrase] = code
			code++
			currentChar = string(c)
		}
	}
	if currentChar != "" {
		result = append(result, mmap[currentChar])
	}
	return result
}
