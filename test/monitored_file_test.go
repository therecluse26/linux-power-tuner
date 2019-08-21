package main

import (
	"github.com/therecluse26/linux-power-tuner/pkg/events/conditions"
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
		regexQuery := ""
		file := conditions.FileMeta{Path: "test_data/search_simple.txt"}
		match := file.SearchFileValue(conditions.Search{Type: conditions.Regex, Query: regexQuery})
		if match != true {
			t.Errorf("got %ts want %t", match, true)
		}
	})

}