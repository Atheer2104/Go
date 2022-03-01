package main

// takes a string url and returns a bool
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// the func returns a map key is a string and bool is a map
func checkWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[bool]string)
	resultChannel := make(chan result)

	for _, url := range urls {
		// starting a goroutine which makes us call multiple func at the same time
		go func(u string) {
			// send statment to resultChannel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// from resultChannel recive statment
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
