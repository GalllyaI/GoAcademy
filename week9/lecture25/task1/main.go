package main

import (
	"database/sql"
	"encoding/json"

	// "encoding/json"
	"goAcademy/week9/lecture25/task1/repository"
	"goAcademy/week9/lecture25/task1/story"
	"log"
	"net/http"

	// "time"

	_ "modernc.org/sqlite"
)

type Storage interface {
	GetStories() []story.NewsItem
}

// func (b *baseURL) HandleTop10JSON() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		ids := b.getIds()
// 		stories := b.getStories(ids)

// 		JSONresp := HandleHackerResponse{stories}

// 		w.Header().Set("Content-Type", "application/json")
// 		v := json.NewEncoder(w)
// 		v.SetIndent("", "    ")
// 		er := v.Encode(&JSONresp)

// 		if er != nil {
// 			http.Error(w, er.Error(), http.StatusBadRequest)
// 			return
// 		}
// 	}
// }

func GetLatestStories(storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// latest := storage.GetLastStoryTimeStamp()
		// period := time.Now().Add(-time.Hour)
		// var stories []NewsItem

		// if latest.Before(period) {
		// 	b := NewBaseURL("https://hacker-news.firebaseio.com")
		// 	ids := b.getIds()
		// 	stories = b.getStories(ids)
		// 	storage.SaveStories(stories)

		// } else {
		stories := storage.GetStories()
		// }

		JSONresp := story.HandleHackerResponse{TopTen: stories}

		w.Header().Set("Content-Type", "application/json")
		v := json.NewEncoder(w)
		v.SetIndent("", "    ")
		er := v.Encode(&JSONresp)

		if er != nil {
			http.Error(w, er.Error(), http.StatusBadRequest)
			return
		}
	}

}

func main() {

	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	mux := http.NewServeMux()
	repo := repository.NewRepository(db)
	mux.Handle("/api/top", GetLatestStories(repo))
	log.Fatal(http.ListenAndServe(":9000", mux))
}
