package lgowithtests

import (
	"maps"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !maps.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for idx := 0; idx < len(urls); idx++ {
		urls[idx] = "a url"
	}
	b.ResetTimer()

	for idx := 0; idx < b.N; idx++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
