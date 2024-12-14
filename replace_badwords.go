package main

import "strings"

func replace_badwords(body string) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		if strings.EqualFold(strings.ToLower(word), "kerfuffle") || strings.EqualFold(strings.ToLower(word), "sharbert") || strings.EqualFold(strings.ToLower(word), "fornax") {
			words[i] = "****"
		}
	}
	cleaned_body := strings.Join(words, " ")
	return cleaned_body
}
