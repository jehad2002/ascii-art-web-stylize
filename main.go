package main

import (
	"net/http"
	"strings"
)

func horizontalMerge(left, right []string) []string {
	var result []string
	minLength := len(left)
	if len(right) < minLength {
		minLength = len(right)
	}
	for i := 0; i < minLength; i++ {
		mergedString := left[i] + right[i]
		result = append(result, mergedString)
	}
	return result
}

func jehad(text string, s int) string {
	var result []string
	mr := []byte(text)
	var getAscii func(byte) []string
	switch s {
	case 1:
		getAscii = standard
	case 2:
		getAscii = shadow
	case 3:
		getAscii = thinkertoy
	}
	for i := 0; i < len(mr); i++ {
		asciim := getAscii(mr[i])
		if i == 0 {
			result = asciim
		} else {
			for j := 0; j < len(result); j++ {
				result[j] += asciim[j]
			}
		}
	}
	return strings.Join(result, "\n")
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	font := r.FormValue("font")
	s := 1
	if font == "shadow" {
		s = 2
	} else if font == "thinkertoy" {
		s = 3
	}
	lines := strings.Split(strings.TrimSpace(text), "\\n")
	asciiArt := []string{}
	for _, line := range lines {
		asciiArt = append(asciiArt, jehad(line, s))
	}
	// Additional Code to Merge ASCII Art
	mergedASCII := strings.Join(asciiArt, "\n")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(mergedASCII))
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	// HTTP server setup
	http.HandleFunc("/generate", generateHandler)
	http.ListenAndServe(":8888", nil)
}
