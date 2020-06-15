package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Handler))
	defer server.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Error("status should be " + http.StatusText(http.StatusTemporaryRedirect))
	}

	location, err := res.Location()
	if err != nil {
		t.Error(err)
	}
	url := location.String()
	if !strings.HasPrefix(url, "https://mcskin.littleservice.cn") &&
		!strings.HasPrefix(url, "https://littleskin.cn") {
		t.Error("incorrect location redirect")
	}

	if !strings.HasSuffix(res.Header.Get("X-Authlib-Injector-API-Location"), "/api/yggdrasil") {
		t.Error("authlib injector location indicator is not found or invalid")
	}

	res, err = client.Get(server.URL + "/user/closet/list?page=1")
	if err != nil {
		t.Error(err)
	}
	location, err = res.Location()
	if err != nil {
		t.Error(err)
	}
	url = location.String()
	if !strings.HasSuffix(url, "/user/closet/list?page=1") {
		t.Error("full url has not been sent")
	}
}
