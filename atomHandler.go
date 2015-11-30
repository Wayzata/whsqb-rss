package whsqbrss

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/gorilla/feeds"

	"appengine"
	"appengine/datastore"
)

func atomHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	query := datastore.NewQuery("Entry").Ancestor(entryKey(c)).Order("-Time")
	entries := make([]*feeds.AtomEntry, 0, 100)
	query.GetAll(c, &entries)
	updatedTime := time.Now().Format(time.RFC3339)
	if len(entries) > 0 {
		updatedTime = entries[0].Updated
	}
	xml, err := xml.Marshal((&feeds.AtomFeed{
		Author: &feeds.AtomAuthor{
			AtomPerson: feeds.AtomPerson{
				Name: "Wayzata High School Quiz Bowl",
				Email: "wayzataquizbowl@gmail.com",
			},
		},
		Entries: entries,
		Link: &feeds.AtomLink{
			Href: "https://whsqb-rss.appspot.com",
		},
		Title: "Next Level News",
		Updated: updatedTime,
	}).FeedXml())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/atom+xml")
	w.Header().Set("Content-Type", "application/xml") // For debugging
	w.Write(xml)
}
