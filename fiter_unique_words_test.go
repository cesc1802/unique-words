package main_test

import (
	"reflect"
	"testing"
	. "unique_words"
)

func TestFilterEmptyList(t *testing.T) {
	uniqueWords := FilterUniqueWords([]string{})
	if len(uniqueWords) != 0 {
		t.Errorf("expected 0, got %d", len(uniqueWords))
	}
}

func TestFilterUniqueWords(t *testing.T) {
	uniqueWords := []string{"abc", "def", "ghi"}
	result := FilterUniqueWords(uniqueWords)

	if !reflect.DeepEqual(result, uniqueWords) {
		t.Errorf(`expected %v, got %v`, uniqueWords, result)
	}
}

func TestFilterNotUniqueWordList(t *testing.T) {
	words := []string{"abc", "def", "ghi", "abc", "ghi"}
	uniqueWords := FilterUniqueWords(words)

	expected := []string{"abc", "def", "ghi"}

	if !reflect.DeepEqual(uniqueWords, expected) {
		t.Errorf(`expected %v, got %v`, expected, uniqueWords)
	}
}
