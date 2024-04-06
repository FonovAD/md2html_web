package apiserver

import (
	"fmt"
	"net/http"
)

// test handle
func (s *APIserver) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello, %s", r.RemoteAddr)))
	}
}

func (s *APIserver) HandleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is a microservice for translating Markdown into HTML format"))
	}
}

func (s *APIserver) HandleMDFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		if r.Header["Content-Type"] == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - need a text file"))
			return
		}
		if r.Header["Content-Type"][0][0:4] != "text" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - need a text file!"))
			return
		}
		buf := make([]byte, r.ContentLength)
		r.Body.Read(buf)
		fmt.Print(string(buf))
		w.Write(buf)
	}
}
