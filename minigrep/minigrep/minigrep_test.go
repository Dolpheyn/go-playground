package minigrep

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	needle, haystack := "how", "how are you doing\nyo wassup"

	found := Search(&needle, &haystack)
	want := []string{"how are you doing"}
	if !reflect.DeepEqual(found, want) {
		t.Errorf("got %v, want %v",found, want)
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	needle, haystack := "hOW", "how are you doing\nyo wassup"

	found := SearchCaseInsensitive(&needle, &haystack)
	want := []string{"how are you doing"}
	if !reflect.DeepEqual(found, want) {
		t.Errorf("got %v, want %v",found, want)
	}
}
