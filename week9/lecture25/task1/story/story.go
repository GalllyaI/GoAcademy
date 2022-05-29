package story

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type NewsItem struct {
	Title     string    `json:"title,omitempty"`
	Score     int       `json:"score,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type HandleHackerResponse struct {
	TopTen []NewsItem `json:"top_stories,omitempty"`
}

type baseURL struct {
	url string
}

func NewBaseURL(url string) *baseURL {
	return &baseURL{url}
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

func (b *baseURL) GetStories() []NewsItem {
	JSONresp := HandleHackerResponse{}

	var news NewsItem
	var wg sync.WaitGroup

	ids := b.getIds()

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
