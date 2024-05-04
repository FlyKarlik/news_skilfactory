package rss

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/FlyKarlik/news_skilfactory/internal/model"
)

func Test_decodeRss(t *testing.T) {
	rssXML := `
	<rss version="2.0">
		<channel>
			<item>
				<title>Test Title 1</title>
				<description>Test Content 1</description>
				<link>http://example.com/1</link>
				<pubDate>Mon, 02 Jan 2006 15:04:05 GMT</pubDate>
			</item>
			<item>
				<title>Test Title 2</title>
				<description>Test Content 2</description>
				<link>http://example.com/2</link>
				<pubDate>Mon, 02 Jan 2006 15:04:05 GMT</pubDate>
			</item>
		</channel>
	</rss>`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(rssXML))
	}))
	defer server.Close()

	posts, err := decodeRss(server.URL)
	if err != nil {
		t.Fatalf("decodeRss returned unexpected error: %v", err)
	}

	expectedPosts := []model.Post{
		{
			Title:   "Test Title 1",
			Content: "Test Content 1",
			Link:    "http://example.com/1",
			PubTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
		{
			Title:   "Test Title 2",
			Content: "Test Content 2",
			Link:    "http://example.com/2",
			PubTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	}

	if !reflect.DeepEqual(posts, expectedPosts) {
		t.Errorf("decoded posts do not match expected posts.\nExpected: %v\nActual: %v", expectedPosts, posts)
	}
}
