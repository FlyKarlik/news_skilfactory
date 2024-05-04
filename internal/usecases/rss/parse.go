package rss

import (
	"encoding/xml"
	"net/http"
	"strconv"
	"time"

	"github.com/FlyKarlik/news_skilfactory/config"
	"github.com/FlyKarlik/news_skilfactory/internal/model"
)

type Parse struct {
	cfg *config.Config
}

func NewParse(cfg *config.Config) *Parse {
	return &Parse{cfg: cfg}
}

func (r *Parse) ParseURL(url string, posts chan<- []model.Post, errs chan<- error) {
	for {
		post, err := decodeRss(url)
		if err != nil {
			errs <- err
		}
		period, err := strconv.Atoi(r.cfg.UpdateTime)
		if err != nil {
			errs <- err
		}
		posts <- post
		time.Sleep(time.Minute * time.Duration(period))
	}
}

func decodeRss(url string) ([]model.Post, error) {
	var rssPost model.RssPost

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err := xml.NewDecoder(resp.Body).Decode(&rssPost); err != nil {
		return nil, err
	}

	var posts []model.Post
	for _, item := range rssPost.Channel.Item {
		var post model.Post

		post.Title = item.Title
		post.Content = item.Description
		post.Link = item.Link
		pubTime, err := validatePubTime(item.PubDate)
		if err != nil {
			return nil, err
		}

		post.PubTime = *pubTime
		posts = append(posts, post)
	}

	return posts, nil
}

func validatePubTime(pubDate string) (*time.Time, error) {
	t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", pubDate)
	if err != nil {
		time, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", pubDate)
		if err != nil {
			return nil, err
		}
		t = time
	}

	return &t, nil
}
