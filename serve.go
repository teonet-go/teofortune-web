// Copyright 2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Teonews website server application.
// Web-server page module.

package main

import (
	"embed"
	"log"
	"net/http"
	"strings"
	"text/template"

	"golang.org/x/crypto/acme/autocert"
)

type Serve struct {
	*Teonet
	domain    string
	templates *template.Template
}

// Page struct send to HTML template
type Page struct {
	Name string
	Body string
}

//go:embed static tmpl
var f embed.FS

// newServe create Serve object and start http server which process http
// requests and communicate with teonet to get / set page values
func newServe(domain string, addr string, teo *Teonet) (err error) {
	s := &Serve{teo, domain, nil}
	err = s.serve(addr)
	return
}

// serve define handlers and start http server
func (s *Serve) serve(addr string) (err error) {

	// Parse template files
	s.templates = template.Must(
		template.ParseFS(f, "tmpl/*.html"),
	)

	// Dynamic handlers
	http.HandleFunc("/", s.homeHandler)

	// Static files handlers
	http.HandleFunc("/favicon.ico", s.faviconHandler)

	// Run web server
	if len(domain) > 0 {
		// Redirect HTTP to HTTPS
		go func() {
			if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
				log.Fatalf("ListenAndServe error: %v", err)
			}
		}()

		// HTTPS server
		err = http.Serve(autocert.NewListener(domain), nil)
	} else {
		// HTTP server
		err = http.ListenAndServe(addr, nil)
	}
	return
}

// redirectTLS redirect HTTP requests to HTTPs
func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+domain+":443"+r.RequestURI,
		http.StatusMovedPermanently)
}

// renderTemplate render template using Page or Rows structure
func (s *Serve) renderTemplate(w http.ResponseWriter, templateName,
	pageTitle string, p interface{}) {

	// Execute selected in function parameters template
	err := s.templates.ExecuteTemplate(w, templateName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// textToHtml converts text to html
func (s *Serve) textToHtml(txt string) string {
	fortune = strings.Replace(fortune, "\n", "<br>\n", -1)
	fortune = strings.Replace(fortune, "\r", "", -1)
	return txt
}

// homeHandler home page handler
func (s *Serve) homeHandler(w http.ResponseWriter, r *http.Request) {
	title := "Teonet Fortune web-site"
	fortune, _ := s.Fortune()
	fortune = s.textToHtml(fortune)
	p := &Page{title, fortune}
	s.renderTemplate(w, "home", title, p)
}

// faviconHandler favicon handler
func (s *Serve) faviconHandler(w http.ResponseWriter, r *http.Request) {
	file := "/static/img/favicon.ico"
	data, _ := f.ReadFile(file)
	w.Write(data)
}
