package server

import (
	"net/http"
)

import (
	"dummy_ai/internal/env"
)

func Start() {

	staticFiles := map[string]string{
		"/manifest.json": "./static/manifest.json",
		"/robots.txt":    "./static/robots.txt",
		"/sitemap.xml":   "./static/sitemap.xml",
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
