package main

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	container := []string{"a", "b", "c", "d"}
	actual := Combinations(container, 2).([][]string)
	expected := [][]string{
		{"a", "b"},
		{"a", "c"},
		{"a", "d"},
		{"b", "c"},
		{"b", "d"},
		{"c", "d"},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
	}
}

func TestPermutations(t *testing.T) {
	container := []string{"a", "b", "c"}
	actual := Permutations(container, 2)
	expected := [][]string{
		{"a", "b"},
		{"b", "a"},
		{"c", "a"},
		{"a", "c"},
		{"b", "c"},
		{"c", "b"},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
	}
}

func TestCartesianProducts(t *testing.T) {
	containers := [][]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	actual := CartesianProducts(containers)
	expected := [][]string{
		{"a", "c", "e"},
		{"b", "c", "e"},
		{"a", "d", "e"},
		{"b", "d", "e"},
		{"a", "c", "f"},
		{"b", "c", "f"},
		{"a", "d", "f"},
		{"b", "d", "f"},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
	}
}

func TestTwines(t *testing.T) {
	containers := [][]string{
		{"a", "b"},
		{"c", "d"},
	}
	actual := Twines(containers).([][]string)
	expected := [][]string{
		{"a", "b", "c", "d"},
		{"a", "c", "b", "d"},
		{"a", "c", "d", "b"},
		{"c", "a", "b", "d"},
		{"c", "a", "d", "b"},
		{"c", "d", "a", "b"},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
	}
}
