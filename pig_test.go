package main

import "testing"

func TestIs_word_started_with_wovel(t *testing.T) {

	got := Is_word_started_with_wovel("ajhjhjkhkjh")
	if !got {
		t.Errorf("Is_word_started_with_wovel %t ERROR", got)
	}

	got = Is_word_started_with_wovel("jhjhjkhkjh")
	if got {
		t.Errorf("Is_word_started_with_wovel %t ERROR", got)
	}
}

func TestIs_char_vowel(t *testing.T) {

	got := Is_char_vowel([]rune("f")[0])
	if got {
		t.Errorf("is char vowel err")
	}

	got = Is_char_vowel([]rune("a")[0])
	if !got {
		t.Errorf("is char vowel err")
	}
}

func TestEnglishWordToPigLatinWord(t *testing.T) {

	text := "away"
	word := EnglishWordToPigLatinWord(&text)
	if word != "awayay" {
		t.Errorf("small vowel error, must: 'awayay', is: '%s\n'", word)
	}

	text = "Away"
	word = EnglishWordToPigLatinWord(&text)
	if word != "Awayay" {
		t.Errorf("title vowel error, must: 'Awayay', is: '%s\n'", word)
	}

	text = "way"
	word = EnglishWordToPigLatinWord(&text)
	if word != "ayway" {
		t.Errorf("small consonant error, must: 'ayway', is: '%s\n'", word)
	}

	text = " Way!!!"
	word = EnglishWordToPigLatinWord(&text)
	if word != " Ayway!!!" {
		t.Errorf("big consonant error, must: ' Ayway!!!', is: '%s\n'", word)
	}

	text = "-Way.."
	word = EnglishWordToPigLatinWord(&text)
	if word != "-Ayway.." {
		t.Errorf("big consonant error, must: '-Ayway..', is: '%s\n'", word)
	}
	text = "- way:"
	word = EnglishWordToPigLatinWord(&text)
	if word != "- ayway:" {
		t.Errorf("big consonant error, must: '- ayway:', is: '%s\n'", word)
	}

	text = "+way;"
	word = EnglishWordToPigLatinWord(&text)
	if word != "+ayway;" {
		t.Errorf("big consonant error, must: '+ayway;', is: '%s\n'", word)
	}

	text = "WayWay"
	word = EnglishWordToPigLatinWord(&text)
	if word != "AyWayway" {
		t.Errorf("big consonant error, must: 'Aywayway', is: '%s\n'", word)
	}
}
