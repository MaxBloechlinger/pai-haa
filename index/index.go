package index

import (
	"strings"
	"unicode"
)

type Index struct {
	postings map[string][]int
}

func New(docs []string) *Index {
	idx := &Index{postings: make(map[string][]int)}
	for docID, doc := range docs {
		seen := make(map[string]struct{})
		for _, term := range tokenize(doc) {
			if _, ok := seen[term]; ok {
				continue
			}
			idx.postings[term] = append(idx.postings[term], docID)
			seen[term] = struct{}{}
		}
	}
	return idx
}

func (idx *Index) Search(term string) []int {
	normalized := normalize(term)
	if normalized == "" {
		return nil
	}
	return append([]int(nil), idx.postings[normalized]...)
}

func tokenize(text string) []string {
	parts := strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	return parts
}

func normalize(text string) string {
	parts := tokenize(text)
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}
