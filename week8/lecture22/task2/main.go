package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type baseURL struct {
	url string
}

func NewBaseURL(url string) *baseURL {
	return &baseURL{url}
}

type HandleHackerResponse struct {
	TopTen []NewsItem `json:"top_stories,omitempty"`
}

type NewsItem struct {
	Title string `json:"title,omitempty"`
	Score int    `json:"score,omitempty"`
}

func (b *baseURL) getIds() []int {
	var result []int
	url := fmt.Sprintf("%s/v0/topstories.json?print=pretty", b.url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, error := http.DefaultClient.Do(req)

	if error != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&result)

	return result[:10]
}

func (b *baseURL) getStories(ids []int) []NewsItem {
	JSONresp := HandleHackerResponse{}

	var news NewsItem
	var wg sync.WaitGroup

	dataChannel := make(chan NewsItem, len(ids))

	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := ids[i]
			detailURL := fmt.Sprintf("%s/v0/item/%d.json?print=pretty", b.url, id)
			detailRes, err := http.Get(detailURL)

			if err != nil {
				log.Fatal(err)
			}

			json.NewDecoder(detailRes.Body).Decode(&news)
			dataChannel <- news
		}()
	}

	wg.Wait()
	close(dataChannel)

	for el := range dataChannel {
		JSONresp.TopTen = append(JSONresp.TopTen, el)
	}

	return JSONresp.TopTen
}

func (b *baseURL) HandleTop10JSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ids := b.getIds()
		stories := b.getStories(ids)

		JSONresp := HandleHackerResponse{stories}

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

	url := NewBaseURL("https://hacker-news.firebaseio.com")
	mux := http.NewServeMux()
	mux.Handle("/api/top", url.HandleTop10JSON())
	log.Fatal(http.ListenAndServe(":9000", mux))
}
