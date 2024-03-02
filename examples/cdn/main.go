package main

import (
	. "github.com/westcoastcode-se/gohtml"
	"github.com/westcoastcode-se/gohtml/a"
	"io"
	"log"
	"net/http"
	"time"
)

// CDN request
func CDN(path string) chan []byte {
	result := make(chan []byte)

	go func() {
		// HTTP client
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			log.Fatal("Error creating HTTP request: ", err.Error())
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Error making HTTP request: ", err.Error())
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			result <- bodyBytes[:]
		}
		close(result)
	}()
	return result
}

// EmitJavascript by making a server-server request to a CDN
func EmitJavascript(path string) Node {
	return EmitChannel(func() chan Node {
		ch := make(chan Node)
		go func() {
			cdn := CDN(path)
			ch <- Script(
				Bytes(<-cdn),
			)
			close(ch)
		}()
		return ch
	})
}

// EmitCSS by making a server-server request to a CDN
func EmitCSS(path string) Node {
	return EmitChannel(func() chan Node {
		ch := make(chan Node)
		go func() {
			cdn := CDN(path)
			ch <- Style(
				Bytes(<-cdn),
			)
			close(ch)
		}()
		return ch
	})
}

// We are using a simple in-memory cache storage. This can be a session-associated cache storage instead, if
// you want to cache html based on the actual logged-in user
var cacheStorage = CreateInMemoryCacheStorage()

func index(w http.ResponseWriter, r *http.Request) {
	// generate the actual html
	_, _ = Html(a.Lang("en"),
		Head(
			// Add a meta header tag with the attribute charset="UTF-8"
			Meta(a.Charset("UTF-8")),
			Title("My Title"),
			// Cache the Materialize CSS in-memory on the server and serve the result backed into the response HTML
			Cache(cacheStorage, "materialize_css", 10*time.Minute,
				EmitCSS("https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css"),
			),
			// Cache the Materialize Javascript in-memory on the server and serve the result backed into the response HTML
			Cache(cacheStorage, "materialize_js", 10*time.Minute,
				EmitJavascript("https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"),
			),
		),
		Body(
			H1(
				Text("This is an example on how to make server-side calls to a CDN and cache the result"),
			),
		),
	)(w)
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
