package repository

import (
	"database/sql"
	"fmt"
	"goAcademy/week9/lecture25/task1/story"
	"log"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (rp *Repository) getLastStoryTimeStamp() time.Time {
	query := "SELECT s.timestamp FROM stories s ORDER BY s.timestamp DESC LIMIT 1"
	var tmstamp time.Time
	if err := rp.db.QueryRow(query).Scan(&tmstamp); err != nil {
		log.Print(err)
	}
	return tmstamp
}

func (rp *Repository) GetStories() []story.NewsItem {

	latest := rp.getLastStoryTimeStamp()
	period := time.Now().Add(-time.Hour)
	// period := time.Now().Add(-2 * time.Second)
	var stories []story.NewsItem

	fmt.Println(latest)
	fmt.Println(period)
	fmt.Println(latest.Before(period))

	if latest.Before(period) {
		b := story.NewBaseURL("https://hacker-news.firebaseio.com")
		stories = b.GetStories()
		rp.deleteStories()
		rp.saveStories(stories)

	}

	query := "SELECT title, score FROM stories"
	rows, err := rp.db.Query(query)

	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		st := story.NewsItem{}
		if err := rows.Scan(&st.Title, &st.Score); err != nil {
			log.Print(err)
		}
		stories = append(stories, st)
	}

	return stories
}

func (rp *Repository) saveStories(sList []story.NewsItem) {
	insertQuery := "INSERT INTO stories(title, score) values(?, ?)"
	for _, s := range sList {
		rp.db.Exec(insertQuery, s.Title, s.Score)
	}
}

func (rp *Repository) deleteStories() {
	res, err := rp.db.Exec("DELETE FROM stories")
	fmt.Println(err, "errror in delete")
	fmt.Println(res.RowsAffected())
}
