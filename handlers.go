package main

import (
	"log"
	"net/http"
	"time"
)

func getTodaysBest(w http.ResponseWriter, r *http.Request) {
	var v viewData
	v.Unpublished = unpublished
	v.MainImg = unpublished[0].MainImg
	exeTmpl(w, r, &v, "home.tmpl")
}
func getUnpublished(w http.ResponseWriter, r *http.Request)       {}
func getRecentlyPublished(w http.ResponseWriter, r *http.Request) {}
func getMostShared(w http.ResponseWriter, r *http.Request)        {}

func create(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	articleTitle := q.Get("t")
	articleText := q.Get("p")
	articleAuthor := q.Get("a")
	connectionIP := r.RemoteAddr
	ID := genPostID(15)
	mainImg := q.Get("i1")
	img2 := q.Get("i2")
	img3 := q.Get("i3")
	var v viewData = viewData{
		MainImg: mainImg,
		Article: article{
			ArticleTitle:  articleTitle,
			ArticleText:   articleText,
			ArticleAuthor: articleAuthor,
			ID:            ID,
			DatePublished: time.Now().Format(time.UnixDate),
			ConnectionIP:  connectionIP,
			MainImg:       mainImg,
			Img2:          img2,
			Img3:          img3,
		},
	}
	addToUnpublished(&v.Article, r.URL)
	exeTmpl(w, r, &v, "create.tmpl")
}

func articleView(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	articleID := q.Get("ID")
	connectionIP := r.RemoteAddr
	log.Println(connectionIP)
	var v viewData
	// v.TickerText = `Ronald Dion DeSantis is an American politician
	// serving since 2019 as the 46th governor of Florida. A member of
	// the Republican Party, DeSantis represented Florida's 6th
	// congressional district in the U.S. House of Representatives
	// from 2013 to 2018`
	v.Article = *unpublished[idToIndex[articleID]]
	v.MainImg = v.Article.MainImg
	v.PageName = "Article"
	exeTmpl(w, r, &v, "article.tmpl")
}
