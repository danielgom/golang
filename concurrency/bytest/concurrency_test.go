package bytest

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type baseSuite struct {
	suite.Suite
}

func TestBaseSuite(t *testing.T) {
	suite.Run(t, &baseSuite{})
}

func (b *baseSuite) mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func (b *baseSuite) Test_checkWebsites() {
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

	got := checkWebsites(b.mockWebsiteChecker, websites)
	b.Equalf(want, got, "check website")
}

// Benchmark testing
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(200 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	numberOfUrls := 100
	urls := make([]string, numberOfUrls)
	for i := range numberOfUrls {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checkWebsites(slowStubWebsiteChecker, urls)
	}
}
