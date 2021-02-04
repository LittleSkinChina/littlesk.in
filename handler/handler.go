package handler

import (
	"net/http"

	"github.com/ip2location/ip2location-go"
)

const (
	domainNameGlobal   string = "https://littleskin.cn"
	domainNameMainland string = "https://mcskin.littleservice.cn"

	dbPath string = "./IP2LOCATION-LITE-DB1.IPV6.BIN"
)

// Handler handles request.
func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := ip2location.OpenDB(dbPath)
	if err != nil {
		makeResponse(domainNameMainland, w, r)
		return
	}

	result, err := db.Get_all(r.RemoteAddr)
	if err != nil {
		makeResponse(domainNameMainland, w, r)
		return
	}
	headers := w.Header()
	// for debugging
	headers.Add("X-IP-Location", result.Country_short)

	if result.Country_short == "CN" {
		makeResponse(domainNameMainland, w, r)
	} else {
		makeResponse(domainNameGlobal, w, r)
	}
}

func makeResponse(domainName string, w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Location", domainName+r.URL.String())
	headers.Add("X-Authlib-Injector-API-Location", domainName+"/api/yggdrasil")

	w.WriteHeader(http.StatusTemporaryRedirect)
}
