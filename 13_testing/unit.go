package unit

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

// named result parameters! can omit result params! good!
func DownloadHTML(url string) (str string, err error) {
	rsp, err := http.Get(url)
	if err != nil {
		return
	}

	bodyReader, err := charset.NewReader(rsp.Body, rsp.Header.Get("Content-Type"))
	if err != nil {
		return
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return
	}
	str = string(body)

	return
}
