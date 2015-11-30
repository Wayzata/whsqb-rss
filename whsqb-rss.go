package whsqbrss

import "net/http"

func init() {
	http.HandleFunc("/atom.xml", atomHandler)
	// http.HandleFunc("/admin", adminHandler)
}
