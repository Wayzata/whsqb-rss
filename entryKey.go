package whsqbrss

import (
	"appengine"
	"appengine/datastore"
)

func entryKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Item", "default", 0, nil)
}
