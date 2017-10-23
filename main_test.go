package main

import "testing"

func TestParseResponse(t *testing.T) {
  json := []byte(`{"message": "bla,bla"}`)
  result := ParseResponse(json)
  if result != "bla,bla" {
    t.Errorf(`Result should be equal "bla,bla", got "%s"`, result)
  }
}
