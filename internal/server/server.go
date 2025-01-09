package server

import (
	"net/http"
)

import (
	"dummy_ai/internal/env"
)

func Start() {

	staticFiles := map[string]string{
		"/manifest.json":            "./static/manifest.json",
		"/robots.txt":               "./static/robots.txt",
		"/sitemap.xml":              "./static/sitemap.xml",
		"/apple_touch_icon.png":     "./static/img/favicon/apple_touch_icon.png",
		"/favicon.ico":              "./static/img/favicon/favicon.ico",
		"/favicon.svg":              "./static/img/favicon/favicon.svg",
		"/favicon_192.png":          "./static/img/favicon/favicon_192.png",
		"/favicon_512.png":          "./static/img/favicon/favicon_512.png",
		"/favicon_512_maskable.png": "./static/img/favicon/favicon_512_maskable.png",
		"/logo.svg":                 "./static/img/logo/logo.svg",
		"/logo_white.svg":           "./static/img/logo/logo_white.svg",
	}

	for route, file := range staticFiles {

		serveFile(route, file)
	}

	if err := http.ListenAndServe(env.ServerAddress(), nil); err != nil {

		panic(err)
	}
}

func serveFile(route string, file string) {

	http.HandleFunc(route, func(responseWriter http.ResponseWriter, request *http.Request) {

		http.ServeFile(responseWriter, request, file)
	})
}
