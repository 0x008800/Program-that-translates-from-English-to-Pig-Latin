//TODO concat strings
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	Vowels string = "aeiouyAEIOUY"
	//	Consonants string = "bcdfghjklmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ"
	Big string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func eng_to_pigLatin() error {

	if len(os.Args) != 2 {
		fmt.Println("Usage: './english_to_pig_latin path_to_your_filename'")
		fmt.Printf("you took %d arguments\n", len(os.Args))
		os.Exit(1)
	}
	//открываем файл переданный в аргументе
	input_file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print("err open inp file\n")
		return err
	}
	defer input_file.Close()

	//создаём файл выхода
	//os.Remove(Add_pig_latin_to_filename(os.Args[1]))
	output_file, err := os.OpenFile(Add_pig_latin_to_filename(os.Args[1]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Print("err open out file\n")
		return err
	}
	defer output_file.Close()
	//fmt.Println(output_file.Name())

	scanner := bufio.NewScanner(input_file)
	writer := bufio.NewWriter(output_file)

	//считываем строки, модифицируем их
	//и записываем в новый файл
	for scanner.Scan() {
		original_string := scanner.Text()
		pig_string := Make_pig_string(original_string)
		_, err := writer.WriteString(pig_string)
		if err != nil {
			fmt.Println("err writing")
		}
		writer.Flush()

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("sc has err")
	}
	return nil
}

//принимает обычную строку слов
//возвращает уже преобразованную в pig-latin
func Make_pig_string(str string) string {

	pig_string := ""
	//разбиваем на слова
	words := strings.Split(str, " ")
	for _, word := range words {
		//каждое слово модифицируем
		//и собираем новую строку
		pig_string += EnglishWordToPigLatinWord(word) + " "
	}
	return pig_string + "\n"
}

//Преобразует слово из английского в поросячью латынь
//*получает слова вместе с точками, запятыми итд
func EnglishWordToPigLatinWord(word string) string {

	//сюда помещаем знаки припинания, которые оказались перед словом
	prefix := ""

	//сюда, те что, после слова
	postfix := ""

	//если в слове первая буква заглавная, нам нужна переменная
	// которая будет помнить это после преобразования
	has_first_symbol_upper := false

	//извлекаем только буквы из полученного "слова"
	reg, err := regexp.Compile("[a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	pure_word := reg.FindString(word) //извлекли

	//если мы вдруг получили отдельностоящие знаки припинания
	// возвращаем их в том же виде
	if len(pure_word) == 0 {
		return word
	}

	//если длина "чистого" слова равна длине полученного
	// значит нам не нужен этот блок
	//и мы пропускаем обработку знаков пунктуации
	if len(word) != len(pure_word) {
		//но если мы здесь
		//извлекаем то что перед и после слова
		tmp := strings.Split(word, pure_word)

		// если извлечённых двое
		if len(tmp) == 2 {
			//отлично, записываем их в переменные
			prefix = tmp[0]
			postfix = tmp[1]
		} else if len(tmp) == 1 { // если же одно
			//нужно выяснить, что это именно

			if strings.HasPrefix(pure_word, word) {

				//"чистое слово" совпадает с началом переданного
				//значит записываем в "после"
				postfix = tmp[0]
			} else {
				//записываем в "до"
				prefix = tmp[0]
			}
		} else {
			panic("reach unreachable")
		}
	}

	//теперь узнаём, у слова первая буква заглавная?
	for _, ch := range Big {
		if strings.HasPrefix(pure_word, string(ch)) {
			//если да - изменяем состояние переменной
			has_first_symbol_upper = true
			break
		}
	}

	// слово начинается с...
	if Is_word_started_with_wovel(pure_word) {
		//гласной
		pure_word += "ay"
		return prefix + pure_word + postfix
	} else {

		//согласной
		part1 := ""
		for _, ch := range pure_word {
			if !Is_char_vowel(ch) {
				part1 += string(ch)
			} else {
				break
			}
		}

		if has_first_symbol_upper {
			//если начиналось с большой буквы
			// и вернуть надо с первой большой
			return prefix + strings.Title(strings.Replace(pure_word, part1, "", 1)+strings.ToLower(part1)+"ay") + postfix
		} else {
			//точно тоже верно и для первых малых букв
			return prefix + strings.Replace(pure_word, part1, "", 1) + strings.ToLower(part1) + "ay" + postfix

		}
	}
}

//предикат "слово начинается с гласной?"
func Is_word_started_with_wovel(word string) bool {
	for _, vowel := range Vowels {
		if strings.HasPrefix(word, string(vowel)) {
			return true
		}
	}
	return false
}

//предикат "первая буква переданного слова гласная?"
func Is_char_vowel(ch rune) bool {

	for _, vowel := range Vowels {

		if vowel == ch {
			return true
		}

	}
	return false
}

//добабляет "(pig-latin)" к сделанному файлу в название перед расширением
// учитывая что в имени файла могут быть несколько точек
func Add_pig_latin_to_filename(s string) string {

	str := strings.Split(s, ".")
	tmp := strings.Join(str[0:len(str)-1], ".") //
	tmp += "(pig-latin)." + str[len(str)-1]
	return tmp
}

//for debug
func check() {
	if len(Vowels) != 12 {
		log.Fatal("wrong count of vowels")
	}

	if len(Big) != 26 {
		log.Fatal("wrong count of Big")
	}

}

func main() {

	//check()

	err := eng_to_pigLatin()
	if err != nil {
		log.Fatal("\tSomething went wrong!!!")
	}
	fmt.Println("\tSUCCESS!!!\n\tfile had been translated to pig-latin!!!")
}
