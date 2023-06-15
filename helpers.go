package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// setupLogging sets output flags and the file for logging
func setupLogging() (f *os.File) {
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(f)

	return
}

// serverFromConf returns a *http.Server with a pre-defined configuration
func serverFromConf() *http.Server {
	return &http.Server{
		Addr:              servicePort,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       5 * time.Second,
	}
}

// exeTmpl is used to build and execute an html template.
func exeTmpl(w http.ResponseWriter, r *http.Request, view *viewData, tmpl string) {
	view.WebsiteName = websiteName
	view.CSS = css
	view.JS = js
	if len(view.TickerText) < 5 {
		view.TickerText = strings.Join(tickers, ` <div class="tickerspacer"></div> `)
	}

	err := templates.ExecuteTemplate(w, tmpl, view)
	if err != nil {
		log.Println(err)
	}
}

// ajaxResponse is used to respond to ajax requests with arbitrary data in the
// format of map[string]string
func ajaxResponse(w http.ResponseWriter, res map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}

// genPostID generates a post ID
func genPostID(length int) (ID string) {
	symbols := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i <= length; i++ {
		s := rand.Intn(len(symbols))
		ID += symbols[s : s+1]
	}
	return
}

// makeZmem returns a redis Z member for use in a ZSET. Score is set to zero
// func makeZmem(st string) redis.Z {
// 	return redis.Z{
// 		Member: st,
// 		Score:  0,
// 	}
// }

func addToUnpublished(a *article, url *url.URL) {
	unpublished = append(unpublished, a)
	idToIndex[a.ID] = len(unpublished)

	writeDB()
}

func readDB() {
	err := json.Unmarshal([]byte(unpub_json), &unpublished)
	if err != nil {
		log.Println(err)
	}
	for k, a := range unpublished {
		idToIndex[a.ID] = k
		tickers = append(tickers, a.TickerText)
	}
}

func writeDB() {
	b, err := json.Marshal(&unpublished)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile("unpublished.json", b, 0644)
	if err != nil {
		log.Println(err)
	}

}
