package main

import "net/http"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	content := "There's an API here"
	replyTextContent(w, req, http.StatusOK, content)
}
