package bytest

type WebsiteChecker func(string) bool

type result struct {
	res  string
	pass bool
}

func checkWebsites(wb WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wb(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.res] = r.pass
	}

	return results
}
