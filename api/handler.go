package handler

import (
	"net/http"

	"github.com/ip2location/ip2location-go"
)

// Handler handles request.
func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := ip2location.OpenDB("./IP2LOCATION-LITE-DB1.IPV6.BIN")
	if err != nil {
		http.Error(w, "failed to open database", http.StatusInternalServerError)
		return
	}

	result, err := db.Get_all(r.RemoteAddr)
	if err != nil {
		http.Error(w, "failed to look up", http.StatusInternalServerError)
		return
	}
	headers := w.Header()
	// for debugging
	headers.Add("X-IP-Location", result.Country_short)

	domainName := ""
	if result.Country_short == "CN" {
		domainName = "https://mcskin.littleservice.cn"
	} else {
		domainName = "https://littleskin.cn"
	}

	headers.Set("Location", domainName+r.URL.String())
	headers.Add("X-Authlib-Injector-API-Location", domainName+"/api/yggdrasil")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
