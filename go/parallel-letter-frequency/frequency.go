package letter

//FreqMap is the map representing text character frequiecies
type FreqMap map[rune]int

//Frequency is the synchronous counting of letter frequency in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// func ConcurrentFrequency(texts []string) FreqMap {
// 	freqChan := make(chan FreqMap, len(texts))
// 	for _, text := range texts {
// 		go func(words string) {
// 			freqChan <- Frequency(words)
// 		}(text)
// 	}
//
// 	result := FreqMap{}
// 	for range texts {
// 		for word, count := range <-freqChan {
// 			result[word] += count
// 		}
// 	}
// 	return result
// }

//ConcurrentFrequency is the asynchronous map of letter frequencies
func ConcurrentFrequency(ss []string) FreqMap {
	out := make(chan FreqMap)
	mainMap := FreqMap{}

	for _, s := range ss {
		go analyze(s, out)
	}

	//Needs to be synchronous
	for i := 0; i < len(ss); i++ {
		m := <-out
		mainMap = consolidateDict(m, mainMap)
	}

	return mainMap
}

func analyze(s string, out chan<- FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	out <- m
}

func consolidateDict(temp, main FreqMap) FreqMap {
	for letter, freq := range temp {
		if f, ok := main[letter]; ok {
			main[letter] = f + freq
		} else {
			main[letter] = freq
		}
	}

	return main
}
