package main

import (
	"github.com/therecluse26/uranium/pkg/events/conditions"
	"testing"
)

func TestSearchCommandResult (t *testing.T) {

	t.Run("simple_search", func(t *testing.T){
		searchValue := 1
		command := conditions.CommandMeta{CommandString: "echo 1"}
		match := command.SearchCommandResult(conditions.Search{Type: conditions.Simple, Query: searchValue})
		if match != true {
			t.Errorf("got %t want %t", match, true)
		}
	})

	t.Run("regex_search_simple", func(t *testing.T){
		regexQuery := `search_simple.*`
		command := conditions.CommandMeta{CommandString: "ls test_data"}
		match := command.SearchCommandResult(conditions.Search{Type: conditions.Regex, Query: regexQuery})
		if match != true {
			t.Errorf("got %ts want %t", match, true)
		}
	})


}