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

	word := EnglishWordToPigLatinWord("away")
	if word != "awayay" {
		t.Errorf("small vowel error, must: 'awayay', is: '%s\n'", word)
	}

	word = EnglishWordToPigLatinWord("Away")
	if word != "Awayay" {
		t.Errorf("title vowel error, must: 'Awayay', is: '%s\n'", word)
	}

	word = EnglishWordToPigLatinWord("way")
	if word != "ayway" {
		t.Errorf("small consonant error, must: 'ayway', is: '%s\n'", word)
	}

	word = EnglishWordToPigLatinWord(" Way!!!")
	if word != " Ayway!!!" {
		t.Errorf("big consonant error, must: ' Ayway!!!', is: '%s\n'", word)
	}

	word = EnglishWordToPigLatinWord("-Way..")
	if word != "-Ayway.." {
		t.Errorf("big consonant error, must: '-Ayway..', is: '%s\n'", word)
	}
	word = EnglishWordToPigLatinWord("- way:")
	if word != "- ayway:" {
		t.Errorf("big consonant error, must: '- ayway:', is: '%s\n'", word)
	}
	word = EnglishWordToPigLatinWord("+way;")
	if word != "+ayway;" {
		t.Errorf("big consonant error, must: '+ayway;', is: '%s\n'", word)
	}
	word = EnglishWordToPigLatinWord("WayWay")
	if word != "AyWayway" {
		t.Errorf("big consonant error, must: 'Aywayway', is: '%s\n'", word)
	}
}
