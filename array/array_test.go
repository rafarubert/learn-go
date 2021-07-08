package main

import "testing"

func TestList(t *testing.T) {
	list := [5]int{1, 2, 3, 4, 5}

	got := List()
	want := list

	if got != want {
		t.Errorf("got %v want %v given, %v", got, want, list)
	}
}
