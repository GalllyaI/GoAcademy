package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetIds(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		json.NewEncoder(w).Encode(&res)
	}))

	client := NewBaseURL(mockServer.URL)
	actual := client.getIds()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestGetStories(t *testing.T) {
	expected := []NewsItem{
		{Title: "top1", Score: 10},
		{Title: "top2", Score: 20},
		{Title: "top3", Score: 30},
		{Title: "top4", Score: 40},
		{Title: "top5", Score: 50},
		{Title: "top6", Score: 60},
		{Title: "top7", Score: 70},
		{Title: "top8", Score: 80},
		{Title: "top9", Score: 90},
		{Title: "top10", Score: 100},
	}

	i := 1
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		title := fmt.Sprintf("top%d", i)
		res := NewsItem{Title: title, Score: i * 10}
		i++

		json.NewEncoder(w).Encode(&res)
	}))

	client := NewBaseURL(mockServer.URL)
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := client.getStories(ids)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}
