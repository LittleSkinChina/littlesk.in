package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
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
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Error("status should be " + http.StatusText(http.StatusTemporaryRedirect))
	}

	location, err := res.Location()
	if err != nil {
		t.Fatal(err)
	}
	url := location.String()
	if !strings.HasPrefix(url, "https://mcskin.littleservice.cn") &&
		!strings.HasPrefix(url, "https://littleskin.cn") {
		t.Error("incorrect location redirect")
	}

	if !strings.HasSuffix(res.Header.Get("X-Authlib-Injector-API-Location"), "/api/yggdrasil") {
		t.Error("authlib injector location indicator is not found or invalid")
	}
}

func TestHandlerWithFullPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Handler))
	defer server.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Get(server.URL + "/user/closet/list?page=1")
	if err != nil {
		t.Fatal(err)
	}
	location, err := res.Location()
	if err != nil {
		t.Fatal(err)
	}
	url := location.String()
	if !strings.HasSuffix(url, "/user/closet/list?page=1") {
		t.Error("full url has not been sent")
	}
}

func TestHandlerWithFailedDB(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Handler))
	defer server.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	os.Rename(dbPath, "./temp.bin")
	defer os.Rename("./temp.bin", dbPath)

	res, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Error("status should be " + http.StatusText(http.StatusTemporaryRedirect))
	}

	location, err := res.Location()
	if err != nil {
		t.Fatal(err)
	}
	if location.String() != "https://mcskin.littleservice.cn/" {
		t.Error("incorrect location redirect")
	}
}
