package server

import (
	"net/http"
	"text/template"
)

import (
	"dummy_ai/internal/env"
)

var (
	serverTemplate = template.Must(template.ParseFiles("./static/server/server.html"))
)

func Start() {

	staticFiles := map[string]string{
		"/manifest.json":            "./static/config/manifest.json",
		"/robots.txt":               "./static/config/robots.txt",
		"/sitemap.xml":              "./static/config/sitemap.xml",
		"/apple_touch_icon.png":     "./static/img/favicon/apple_touch_icon.png",
		"/favicon.ico":              "./static/img/favicon/favicon.ico",
		"/favicon.svg":              "./static/img/favicon/favicon.svg",
		"/favicon_192.png":          "./static/img/favicon/favicon_192.png",
		"/favicon_512.png":          "./static/img/favicon/favicon_512.png",
		"/favicon_512_maskable.png": "./static/img/favicon/favicon_512_maskable.png",
		"/logo.svg":                 "./static/img/logo/logo.svg",
		"/logo_white.svg":           "./static/img/logo/logo_white.svg",
		"/carbon.css":               "./static/lib/carbon/carbon.css",
		"/carbon.js":                "./static/lib/carbon/carbon.js",
		"/wasm_exec.js":             "./static/lib/wasm/wasm_exec.js",
		"/wasm_start.js":            "./static/lib/wasm/wasm_start.js",
		"/error_404.wasm":           "./static/wasm/error_404.wasm",
		"/index.wasm":               "./static/wasm/index.wasm",
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

func error404(responseWriter http.ResponseWriter) {

	responseWriter.WriteHeader(http.StatusNotFound)
	executeServerTemplate(responseWriter, "/error_404.wasm")
}

func executeServerTemplate(responseWriter http.ResponseWriter, wasmRoute string) {

	if err := serverTemplate.Execute(responseWriter, wasmRoute); err != nil {

		panic(err)
	}
}
