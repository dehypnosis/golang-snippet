package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -run="none" -bench . -benchmem
func BenchmarkDownloadHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DownloadHTML("http://google.com")
	}
}

// go test -run ^TestDownload -cover -v
func ExampleDownloadHTML() {
	DownloadHTML("http://google.com")
	fmt.Printf("output check")
	// Output: output check
}
func TestDownloadSingleHTML(t *testing.T) {

	url := "https://www.google.co.kr/search?q=테스트"
	if _, err := DownloadHTML(url); err != nil {
		t.Errorf("%s\nFailed to download 😱\n, %v", url, err)
	}
	t.Logf("%s\nDownloaded successfully 😀\n", url)

}

func TestDownloadMultipleHTMLs(t *testing.T) {
	urls := []string{
		"https://naver.com",
		"https://benzen.io",
	}
	for index, url := range urls {
		t.Run(fmt.Sprintf("ex%d", index), func(t *testing.T) {
			if _, err := DownloadHTML(url); err != nil {
				t.Fatalf("%s\nFailed to download 😱\n, %v", url, err)
			}
			t.Logf("%s\nDownloaded successfully 😀\n", url)
		})
	}
}

// mock http server
func createMockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello world!"))
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	println("new mock server", server.URL)
	return server
}

var httpServer *httptest.Server

const testEndpointAnswer = "Hello Again!"

func init() {
	// run mock server
	httpServer = createMockServer()
	//defer httpServer.Close()

	// route some enpoint on default http server
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(testEndpointAnswer))
	})
}

func TestDownloadSingleHTMLFromMockServer(t *testing.T) {
	url := httpServer.URL

	if _, err := DownloadHTML(url); err != nil {
		t.Errorf("%s\nFailed to download 😱\n, %v", url, err)
		return
	}
	t.Logf("%s\nDownloaded successfully 😀\n", url)
}

func TestDownloadSingleEndpoint(t *testing.T) {
	// using (http.DefaultServeMux or any http server instance).ServeHTTP to not to boot server

	rs := httptest.NewRecorder()
	rq, _ := http.NewRequest("GET", "http://host_is_not_important_for_single_server.com/test", nil)
	http.DefaultServeMux.ServeHTTP(rs, rq)

	body := rs.Body.String()
	t.Logf("Response is %s", body)

	if body != "Hello Again!" {
		t.Errorf("Response is not equal to %s", testEndpointAnswer)
	}
}
