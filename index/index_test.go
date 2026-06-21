package index

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	idx := New([]string{
		"Go makes search engines",
		"Search with inverted index in Go",
		"Crawlers feed the index",
	})

	if got := idx.Search("go"); !reflect.DeepEqual(got, []int{0, 1}) {
		t.Fatalf("go: got %v", got)
	}
	if got := idx.Search("index"); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("index: got %v", got)
	}
	if got := idx.Search("missing"); len(got) != 0 {
		t.Fatalf("missing: got %v", got)
	}
}
