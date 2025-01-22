package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type viewData struct {
	WebsiteName string
	Article     article
	Articles    []*article
	Unpublished []*article
	CSS         string
	JS          string
	TickerText  string
	MainImg     string
	PageName    string
}

type article struct {
	ArticleText   string `json:"articleText"`
	ArticleTitle  string `json:"articleTitle"`
	ArticleAuthor string `json:"articleAuthor"`
	DatePublished string `json:"datePublished"`
	MainImg       string `json:"mainImg"`
	Img2          string `json:"img2"`
	Img3          string `json:"img3"`
	ConnectionIP  string `json:"connectionIP"`
	TickerText    string `json:"tickerText"`
	UniqueViews   int    `json:"uniqueViews"`
	ID            string `json:"ID"`
}

// ckey/ctxkey is used as the key for the HTML context and is how we retrieve
// token information and pass it around to handlers
type ckey int

const ctxkey ckey = iota

var (
	//go:embed public/css/main.css
	css string
	//go:embed public/js/main.js
	js string
	//go:embed final.json
	unpub_json string
	//go:embed published.json
	pub_json    string
	servicePort = ":10528"
	logFilePath = "log.txt"

	templates          = template.Must(template.New("main").ParseGlob("internal/pages/*"))
	mux                = http.NewServeMux()
	websiteName string = "[ <i>(the) U N S O U R C E D</i> ] • The leading source in <u><i>unsourced</i></u> news"
	published   []*article
	unpublished []*article
	// idToLink    map[string]string = make(map[string]string)
	idToIndex map[string]int = make(map[string]int)
	tickers   []string
)

func main() {
	logFile := setupLogging()
	defer logFile.Close()

	template.Must(templates.ParseGlob("internal/components/*/*"))
	mux.HandleFunc("/", getTodaysBest)
	mux.HandleFunc("/review", getUnpublished)
	mux.HandleFunc("/shiny", getRecentlyPublished)
	mux.HandleFunc("/gold", getMostShared)

	mux.HandleFunc("/article/", articleView)
	mux.HandleFunc("/create/", create)
	mux.HandleFunc("/about", articleView)

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	srv := serverFromConf()
	ctx, cancelCtx := context.WithCancel(context.Background())

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
		cancelCtx()
	}()

	readDB()
	fmt.Println("Server started @ http://localhost" + srv.Addr)
	log.Println("Server started @ " + srv.Addr)
	<-ctx.Done()
	log.Println("exited")
}
