package main

import (
	"github.com/therecluse26/uranium/pkg/events/conditions"
	"testing"
)

func TestSearchFileValue (t *testing.T) {

	t.Run("simple_search", func(t *testing.T){
		searchValue := 1
		file := conditions.FileMeta{Path: "test_data/search_simple.txt"}
		match := file.SearchFileValue(conditions.Search{Type: conditions.Simple, Query: searchValue})
		if match != true {
			t.Errorf("got %t want %t", match, true)
		}
	})

	t.Run("regex_search_simple", func(t *testing.T){
		regexQuery := `TRUE.*`
		file := conditions.FileMeta{Path: "test_data/search_regex_simple.txt"}
		match := file.SearchFileValue(conditions.Search{Type: conditions.Regex, Query: regexQuery})
		if match != true {
			t.Errorf("got %ts want %t", match, true)
		}
	})

	t.Run("regex_search_keyval", func(t *testing.T){
		regexQuery := `status=true`
		file := conditions.FileMeta{Path: "test_data/regex_search_keyval.txt"}
		match := file.SearchFileValue(conditions.Search{Type: conditions.Regex, Query: regexQuery})
		if match != true {
			t.Errorf("got %ts want %t", match, true)
		}
	})

}