package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

type HackerData struct {
	PageTitle string
	Links     []NewsItem
}

type HandleHackerResponse struct {
	TopTen []NewsItem `json:"top_stories,omitempty"`
}

type NewsItem struct {
	Title string `json:"title,omitempty"`
	Score int    `json:"score,omitempty"`
}

func HandleTop10JSON() http.HandlerFunc {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, error := http.DefaultClient.Do(req)

		if error != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer res.Body.Close()

		var resp []int
		json.NewDecoder(res.Body).Decode(&resp)
		// w.Header().Set("Content-Type", "application/json")
		JSONresp := HandleHackerResponse{}

		var news NewsItem
		var wg sync.WaitGroup

		for i := 0; i < 10; i++ {
			i := i
			wg.Add(1)
			go func() {
				defer wg.Done()

				id := resp[i]
				detailURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)
				detailRes, err := http.Get(detailURL)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				// w.Header().Set("Content-Type", "application/json")
				json.NewDecoder(detailRes.Body).Decode(&news)
				JSONresp.TopTen = append(JSONresp.TopTen, news)
			}()
		}
		wg.Wait()

		t, err := template.ParseFiles("index.html")

		if err != nil {
			fmt.Println(err)
		}

		er := t.Execute(w, HackerData{PageTitle: "Top ten", Links: JSONresp.TopTen})

		if er != nil {
			fmt.Println(er)
		}
	}
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/api/top", HandleTop10JSON())
	log.Fatal(http.ListenAndServe(":9000", mux))
}
