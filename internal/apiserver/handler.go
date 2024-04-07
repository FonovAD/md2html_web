package apiserver

import (
	"encoding/json"
	"fmt"
	MarkdownToHTML "md2html_web/pkg/md2html"
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
		s.logger.Info(fmt.Sprintf("%d\t%s", http.StatusOK, r.URL.Path))
	}
}

func (s *APIserver) HandleMDFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		if r.Header["Content-Type"] == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - empty body"))
			s.logger.Warnf("%d\t%s", http.StatusBadRequest, r.URL.Path)
			return
		}
		if r.Header["Content-Type"][0][0:4] != "text" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - need a text file!"))
			s.logger.Warnf("%d\t%s", http.StatusBadRequest, r.URL.Path)
			return
		}
		buf := make([]byte, r.ContentLength)
		r.Body.Read(buf)
		html, err := MarkdownToHTML.Convert(string(buf))
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("I can't make an analysis"))
			s.logger.Infof("%d\t%s", http.StatusNoContent, r.URL.Path)
			return
		}
		fmt.Print(html)
		w.Write([]byte(html))
		s.logger.Infof("%d\t%s", http.StatusOK, r.URL.Path)
	}
}

func (s *APIserver) HandleMDBody() http.HandlerFunc {
	type request struct {
		Text string `json:"Text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if r.Header["Content-Type"] == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - empty body"))
			return
		}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			s.logger.Warnf("%d\t%s", http.StatusBadRequest, r.URL.Path)
			return
		}
		html, err := MarkdownToHTML.Convert(string(req.Text))
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("I can't make an analysis"))
			s.logger.Warnf("%d\t%s", http.StatusNoContent, r.URL.Path)
			return
		}
		w.Write([]byte(html))
		s.logger.Infof("%d\t%s", http.StatusOK, r.URL.Path)
	}
}
